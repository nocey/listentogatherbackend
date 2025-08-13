package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/listentogether/auth"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
	logger "github.com/listentogether/log"
)

type User struct {
}

func (u *User) GetAll(c *fiber.Ctx) error {
	var users []models.Users
	err := database.DBConn.Debug().Preload("Permissions").Find(&users).Error
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *User) Login(c *fiber.Ctx) error {
	payload := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Name == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name and password are required",
		})
	}
	user := &models.Users{}
	err := database.DBConn.Debug().Where("name = ?", payload.Name).First(&user).Error
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": fmt.Sprintf("User %s not found", payload.Name),
		})
	}

	if err := auth.GeneratePasswordHash(&payload.Password); err != nil {
		logger.Error("Error generating password hash:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate password hash",
			"error":   err.Error(),
		})
	}

	if user.Password != payload.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid password",
		})
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
	})

	config, _ := config.Load()

	token, err := jwtToken.SignedString(config.JwtToken)
	if err != nil {
		logger.Error("Error signing JWT token:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (u *User) SignUp(c *fiber.Ctx) error {
	payload := struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if payload.Name == "" || payload.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Name and password are required",
		})
	}
	err := auth.GeneratePasswordHash(&payload.Password)
	logger.Debug("Creating user:", payload.Password)

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	user := &models.Users{
		Name:     payload.Name,
		Password: payload.Password,
	}

	err = database.DBConn.Debug().Create(&user).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create user",
			"error":   err.Error(),
		})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}
