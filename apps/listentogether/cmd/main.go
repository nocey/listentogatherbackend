package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/auth"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
	"github.com/listentogether/main/middlewares"
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}
	app := fiber.New()
	database.Connect(&config.Database)

	app.Use("/auth", func(c *fiber.Ctx) error {
		auth.Protected()

		return c.JSON(fiber.Map{"status": fiber.StatusOK, "meesage": "Auth service is working"})
	})
	app.Use(middlewares.UseConfig)

	app.Get("/", func(c *fiber.Ctx) error {
		var user []models.User
		database.DBConn.Model(&models.User{}).Select("*").Find(&user)
		return c.Status(fiber.StatusOK).JSON(user)
	})

	fmt.Println("Starting ListenTogether server on port:", os.Getenv("PORT"))

	log.Fatal(app.Listen(":" + config.Port))
}
