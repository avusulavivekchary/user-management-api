package main

import (
	"log"

	"github.com/avusulavivekchary/user-management-api/config"
	"github.com/avusulavivekchary/user-management-api/db/sqlc"
	"github.com/avusulavivekchary/user-management-api/internal/handler"
	"github.com/avusulavivekchary/user-management-api/internal/logger"
	"github.com/avusulavivekchary/user-management-api/internal/middleware"
	"github.com/avusulavivekchary/user-management-api/internal/repository"
	"github.com/avusulavivekchary/user-management-api/internal/routes"
	"github.com/avusulavivekchary/user-management-api/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = logger.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Log.Sync()

	conn, err := config.ConnectDB()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	logger.Log.Info("Database connected successfully")

	queries := sqlc.New(conn)

	userRepo := repository.NewUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "API is running",
		})
	})

	routes.SetupRoutes(app, userHandler)
	logger.Log.Info("Server running on port 3000")

	log.Fatal(app.Listen(":3000"))
}
