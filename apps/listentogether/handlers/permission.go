package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
)

type Permission struct {
}

func (p *Permission) GetAll(c *fiber.Ctx) error {
	var permissions []models.Permissions
	err := database.DBConn.Find(&permissions).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(permissions)
}

func (p *Permission) GetUserPermissions(c *fiber.Ctx) error {
	var userPermissions []models.Permissions
	user := c.Locals("user").(*models.Users)

	err := database.DBConn.Find(&userPermissions).Where("user_id = ?", user.ID).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(userPermissions)
}
