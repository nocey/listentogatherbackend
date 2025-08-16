package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/database/models"
)

type Config struct {
}

func (config *Config) GetAllConfigs(c *fiber.Ctx) error {
	configs := models.AppConfig{}
	configssList := configs.GetAllWithInObject()
	if configssList == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "no configurations found",
		})
	}
	return c.Status(fiber.StatusOK).JSON(configssList)
}
