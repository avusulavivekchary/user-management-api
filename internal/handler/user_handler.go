package handler

import (
	"strconv"

	"github.com/avusulavivekchary/user-management-api/internal/middleware"
	"github.com/avusulavivekchary/user-management-api/internal/models"
	"github.com/avusulavivekchary/user-management-api/internal/service"
	"github.com/gofiber/fiber/v2"
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

	if err := middleware.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
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
			"age":  service.CalculateAge(user.Dob.Time),
		})
	}

	return c.JSON(response)
}
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	user, err := h.service.GetUserByID(int32(id))
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
		"age":  service.CalculateAge(user.Dob.Time),
	})
}
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	var req models.UpdateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	if err := middleware.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user, err := h.service.UpdateUser(
		int32(id),
		req.Name,
		req.DOB,
	)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Time.Format("2006-01-02"),
	})
}
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	err = h.service.DeleteUser(int32(id))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
