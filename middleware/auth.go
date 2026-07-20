package middleware

import (
	"strings"

	"github.com/LeannXy/Project_Management/config"
	"github.com/LeannXy/Project_Management/utils"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return utils.Unauthorized(c, "Unauthorized", "Token missing")
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return utils.Unauthorized(c, "Unauthorized", "Invalid token")
	}

	return c.Next()
}