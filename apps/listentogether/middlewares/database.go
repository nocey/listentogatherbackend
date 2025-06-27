package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
)

func UseDatabase(c *fiber.Ctx) error {
	config := c.Locals("config").(*config.Config)
	db, err := database.Connect(&config.Database)
	if err != nil {
		return fmt.Errorf("%s: Database connection error", err)
	}

	c.Locals("db", db)
	return c.Next()
}
