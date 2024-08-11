package main

import (
	"api-rest-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	routes.UseRoutes(app)

	app.Listen(":3000")
}
