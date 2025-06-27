package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
)

func GetAll(c *fiber.Ctx) error {
	var users []models.User
	database.DBConn.Model(&models.User{}).Find(&users)

	return c.Status(fiber.StatusOK).JSON(users)
}
