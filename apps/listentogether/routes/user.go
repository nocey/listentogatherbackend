package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/main/handlers/user"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userGroup.Get("/", user.GetAll)
}
