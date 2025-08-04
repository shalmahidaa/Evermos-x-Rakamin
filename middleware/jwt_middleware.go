package middleware

import (
	"evermos/utils"
	"github.com/gofiber/fiber/v2"
	//"github.com/golang-jwt/jwt/v5"
	"strings"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "missing or invalid token"})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		token, claims, err := utils.VerifyJWT(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "invalid token"})
		}

		// Inject ke context
		c.Locals("user_id", uint(claims["user_id"].(float64)))
		c.Locals("role", claims["role"].(string))
		return c.Next()
	}
}

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Ambil role dari context (diset di middleware JWTProtected)
		role := c.Locals("role")
		if role != "admin" {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Admin only",
			})
		}
		return c.Next()
	}
}