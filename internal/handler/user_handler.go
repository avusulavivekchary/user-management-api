package handler

import (
	"github.com/gofiber/fiber/v2"

	"github.com/avusulavivekchary/user-management-api/internal/models"
	"github.com/avusulavivekchary/user-management-api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	user, err := h.service.CreateUser(req.Name, req.DOB)
	if err != nil {
		println(err.Error())

		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
	})
}
func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers()
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	var response []fiber.Map

	for _, user := range users {
		response = append(response, fiber.Map{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Time.Format("2006-01-02"),
		})
	}

	return c.JSON(response)
}
