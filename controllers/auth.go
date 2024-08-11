package controllers

import (
	"api-rest-go/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LoginController(c *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}

	if body.Email == "" || body.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "invalid body"})
	}

	token := utils.SignedToken(body.Email)

	return c.JSON(fiber.Map{"isSuccess": true, "message": "login success", "data": "", "token": token})

}

func LogoutController(c *fiber.Ctx) error {

	token := c.Get("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	claims, err := utils.VerifyToken(token)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "unauthorized"})
	}

	return c.JSON(fiber.Map{"isSuccess": true, "message": "logout success", "data": "", "token": claims})
}
