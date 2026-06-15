package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/avusulavivekchary/user-management-api/internal/handler"
)

func SetupRoutes(app *fiber.App, userHandler *handler.UserHandler) {
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
}
