package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
)

// This checks that incoming requests have valid session tokens
func AuthMiddleware(c *fiber.Ctx) error { // Does session token exist?
	sessionToken := c.Cookies("session_token")
	if sessionToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var user model.User
	if err := model.DB.First(&user, "session_token = ?", sessionToken).Error; err != nil { // Does session token match username?
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid session"})
	}

	c.Locals("user", user) // Store the user with the request
	return c.Next()        // Passes the request to the next stage
}

// This looks overcommented, and it may be, but its because going through this we are all learning
// The comments ensure that everyone working on the project knows what each line does
