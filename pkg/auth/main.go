package auth

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/listentogether/config"
	"github.com/listentogether/database"
	"github.com/listentogether/database/models"
	logger "github.com/listentogether/log"
)

type ClaimUser struct {
	jwt.MapClaims
	UserName string `json:"username"`
}

func Protected(token string) (*models.Users, error) {
	user := &models.Users{}
	if token == "" {
		return nil, fmt.Errorf("missing token")
	}

	parts := strings.Split(token, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, fmt.Errorf("invaild header format")
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
		return nil, fmt.Errorf("jwt parsing error")
	}

	if claims, ok := jwtToken.Claims.(ClaimUser); ok {
		database.DBConn.Model(&user).Where("name = ?", claims.UserName).Find(&user)
	}

	return user, nil
}

func Middleware(perm *models.Permissions) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		user, err := Protected(token)
		if err != nil {
			logger.Debug(err)
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Authocentation is required",
				"error":   err.Error(),
			})
		}
		if perm != nil {
			user.HasPermission(perm)
		}
		c.Locals(user, "user")

		return c.Next()
	}
}

func JWTtoken(user *models.Users) (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", fmt.Errorf("loading config error on auth pkg")
	}

	claims := ClaimUser{
		MapClaims: jwt.MapClaims{
			"exp": jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
		UserName: user.Name,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(cfg.JwtToken)
}

func GeneratePasswordHash(password *string) error {
	if *password == "" {
		return fmt.Errorf("password cannot be empty")
	}
	config, _ := config.Load()
	*password = fmt.Sprintf("%s%s", *password, config.Salt)
	hash := sha256.New()
	hash.Write([]byte(*password))

	*password = fmt.Sprintf("%x", hash.Sum(nil))

	return nil
}
