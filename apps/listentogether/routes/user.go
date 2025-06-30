package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/main/handlers/user"
	"github.com/listentogether/main/middlewares"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userGroup.Get("/", middlewares.Auth(), user.GetAll)
}
