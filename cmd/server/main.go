package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/avusulavivekchary/user-management-api/config"
	"github.com/avusulavivekchary/user-management-api/db/sqlc"
	"github.com/avusulavivekchary/user-management-api/internal/handler"
	"github.com/avusulavivekchary/user-management-api/internal/repository"
	"github.com/avusulavivekchary/user-management-api/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	log.Println("Database connected successfully")

	queries := sqlc.New(conn)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "API is running",
		})
	})

	app.Post("/users", userHandler.CreateUser)
	app.Get("/users", userHandler.GetUsers)
	app.Get("/users/:id", userHandler.GetUserByID)
	app.Put("/users/:id", userHandler.UpdateUser)
	app.Delete("/users/:id", userHandler.DeleteUser)
	log.Println("Server running on port 3000")

	log.Fatal(app.Listen(":3000"))
}
