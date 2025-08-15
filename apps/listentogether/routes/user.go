package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/auth"
	"github.com/listentogether/main/handlers"
)

func UserRoutes(app *fiber.App) {
	userGroup := app.Group("/users")

	userHandler := &handlers.User{}
	permissionHandler := &handlers.Permission{}
	userGroup.Get("/", userHandler.GetAll)
	userGroup.Get("/information", auth.Middleware(""), userHandler.GetInformation)
	userGroup.Get("/permissions", auth.Middleware(""), permissionHandler.GetUserPermissions)
	userGroup.Post("/login", userHandler.Login)
	userGroup.Post("/signup", userHandler.SignUp)
}
