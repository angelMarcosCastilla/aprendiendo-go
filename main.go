package main

import (
	"api-rest-go/db"
	"api-rest-go/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	conexion := db.ConnectDB()

	// insertar un usuario
	_, err := conexion.Exec("insert into usuarios (username, email, password) values ('user1', 'user1@example.com', '123456')")
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World!")
	})

	routes.UseRoutes(app)

	app.Listen(":3000")
}
