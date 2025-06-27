package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/config"
)

func UseConfig(c *fiber.Ctx) error {
	config, err := config.Load()
	if err != nil {
		return fmt.Errorf("Failed to load configuration: %v\n", err)
	}

	c.Locals("config", config)

	return c.Next()
}
