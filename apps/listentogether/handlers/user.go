package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
)

type User struct {
}

func (u *User) GetAll(c *fiber.Ctx) error {
	var users []models.User
	err := database.DBConn.Debug().Preload("Permissions").Find(&users).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}
