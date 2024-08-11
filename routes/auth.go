package routes

import (
	"api-rest-go/controllers"

	"github.com/gofiber/fiber/v2"
)

func useRouteLogin(app *fiber.App) {
	app.Post("/login", controllers.LoginController)
	app.Post("/logout", controllers.LogoutController)
}
