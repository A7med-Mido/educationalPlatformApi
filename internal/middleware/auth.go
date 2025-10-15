package middleware

import (
	"github.com/A7med-Mido/educationalPlatformApi/internal/utils"
	"github.com/A7med-Mido/educationalPlatformApi/internal/utils/constants"
	"github.com/gofiber/fiber/v3"
)

func AuthMiddleware(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(constants.Unauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}
		claims, err := utils.ParseJWT(token)
		if err != nil || claims.Role != requiredRole {
			return c.Status(403).JSON(fiber.Map{"error": "Forbidden"})
		}
		c.Locals("userID", claims.UserID)
		return c.Next()
	}
}