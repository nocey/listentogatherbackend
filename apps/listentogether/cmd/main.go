package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	"github.com/listentogether/main/routes"
)

func main() {
	config, err := config.Load()
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		os.Exit(1)
	}
	app := fiber.New()
	database.Connect(&config.Database)

	routes.UserRoutes(app)

	fmt.Println("Starting ListenTogether server on port:", os.Getenv("PORT"))

	log.Fatal(app.Listen(":" + config.Port))
}
