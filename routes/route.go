package routes

import "github.com/gofiber/fiber/v2"

func UseRoutes(app *fiber.App) {
	useRouteLogin(app)
	UseRouteUser(app)
}
