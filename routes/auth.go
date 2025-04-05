package routes

import (
	"github.com/loganstottle/CashCraft/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", loginHandler)
	app.Post("/logout", logoutHandler)
	app.Get("/me", authMiddleware, meHandler)
}

func loginHandler(c *fiber.Ctx) error {
	type LoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user model.User
	if err := model.DB.First(&user, "email = ?", input.Email).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	if user.Password != model.HashPassword(input.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	sessionToken := uuid.New().String()
	user.SessionToken = sessionToken
	model.DB.Save(&user)

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.JSON(fiber.Map{"message": "Logged in!"})
}

func logoutHandler(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")
	model.DB.Model(&model.User{}).Where("session_token = ?", sessionToken).Update("session_token", nil)

	c.ClearCookie("session_token")
	return c.JSON(fiber.Map{"message": "Logged out"})
}

func meHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
	})
}
