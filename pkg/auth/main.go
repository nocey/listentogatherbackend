package auth

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
)

type ClaimUser struct {
	jwt.MapClaims
	UserName string `json:"username"`
}

func Protected(token string) (*models.User, error) {
	user := &models.User{}
	if token == "" {
		return nil, fmt.Errorf("Missing token")
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("Invaild header format")
	}
	token = parts[1]

	cfg, err := config.Load()
	if err != nil {
		return nil, fmt.Errorf("loading config error on auth pkg")
	}

	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return cfg.JwtToken, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Jwt parsing error")
	}

	if claims, ok := jwtToken.Claims.(ClaimUser); ok {
		database.DBConn.Model(&user).Where("name = ?", claims.UserName).Find(&user)
	}

	return user, nil
}

func Auth(perm *models.Permissions) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		user, err := Protected(token)
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusUnauthorized).Send([]byte("Authocentation is required"))
		}
		if perm != nil {
			user.HasPermission(perm)
		}
		c.Locals(user, "user")

		return c.Next()
	}
}
