package middlewares

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/auth"
)

func Auth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		user, err := auth.Protected(token)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).Send([]byte("Authocentation is required"))
		}
		c.Locals(user, "user")

		return c.Next()
	}
}
