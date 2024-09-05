package handler

import (
	"go-auth/internal/service"

	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json :"password"`
	}
	var req Request

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"Error": "invalid req",
		})
	}
	err := service.RegisterUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User rgister success",
	})

}

func Login(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var req Request
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	token, err := service.AuthenticateUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid credentials",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
