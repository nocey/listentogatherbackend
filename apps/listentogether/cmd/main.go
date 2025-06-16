package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/auth"
)

func main()  {
	app := fiber.New()

	app.Use("/auth", func(c *fiber.Ctx) error {
		auth.Protected()

		return c.JSON(fiber.Map{"status": fiber.StatusOK, "meesage": "Auth service is working"})
	})

	log.Fatal(app.Listen(":3000"))
}