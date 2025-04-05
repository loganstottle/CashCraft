package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")
	if sessionToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user model.User
	if err := model.DB.First(&user, "session_token = ?", sessionToken).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session"})
	}

	c.Locals("user", user)
	return c.Next()
}
