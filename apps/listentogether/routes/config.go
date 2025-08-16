package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/main/handlers"
)

func ConfigRoutes(app *fiber.App) {
	configGroup := app.Group("/config")

	configHandler := &handlers.Config{}
	configGroup.Get("/", configHandler.GetAllConfigs)
}
