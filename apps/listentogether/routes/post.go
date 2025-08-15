package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/auth"
	"github.com/listentogether/main/handlers"
)

func PostRoutes(app *fiber.App) {
	postGroup := app.Group("/posts")

	postHandler := &handlers.Post{}
	postGroup.Post("/", auth.Middleware("create_post"), postHandler.Create)
	postGroup.Get("/:id", auth.Middleware("read_post"), postHandler.Create)
	postGroup.Patch("/:id", auth.Middleware("update_post"), postHandler.Create)
	postGroup.Delete("/:id", auth.Middleware("delete_post"), postHandler.Create)
}
