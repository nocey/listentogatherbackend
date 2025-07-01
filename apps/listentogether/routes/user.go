package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/main/handlers"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userHandler := &handlers.User{}
	userGroup.Get("/", userHandler.GetAll)
}
