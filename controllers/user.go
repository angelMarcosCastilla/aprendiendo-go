package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func CreateUserController(c *fiber.Ctx) error {

	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Username string `json:"username"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "create user success"})
}
