package routes

import (
	"api-rest-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func UseRouteUser(app *fiber.App) {
	app.Post("/user", controllers.CreateUserController)
}
