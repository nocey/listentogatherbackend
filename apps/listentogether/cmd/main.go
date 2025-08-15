package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	logger "github.com/listentogether/log"
	"github.com/listentogether/main/routes"
)

func main() {
	config, err := config.Load()
	logger.New("output.log", logger.DEBUG)
	if err != nil {
		logger.Debug(fmt.Sprintf("Failed to load configuration: %v\n", err))
		os.Exit(1)
	}
	app := fiber.New()
	database.Connect(&config.Database)

	routes.UserRoutes(app)
	routes.PostRoutes(app)

	logger.Debug("Starting ListenTogether server on port:", config.Port)
	logger.Fatal(app.Listen(":" + config.Port))
}
