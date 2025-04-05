package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/register", RegisterHandler)
	app.Post("/login", LoginHandler)
	app.Post("/logout", LogoutHandler)
	app.Get("/me", AuthMiddleware, MeHandler)
}

func RegisterHandler(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user model.User

	if err := model.DB.First(&user, "username = ?", input.Username).Error; err == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Account already exists"})
	}

	user.Username = input.Username
	user.Password = model.HashPassword(input.Password)

	user.Cash = 100000.00

	sessionToken := uuid.New().String()
	user.SessionToken = sessionToken
	model.DB.Create(&user)

	c.Cookie(&fiber.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
	})

	return c.JSON(fiber.Map{"message": "Logged in!"})
}

func LoginHandler(c *fiber.Ctx) error {
	var input LoginInput
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user model.User
	if err := model.DB.First(&user, "username = ?", input.Username).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// fmt.Printf("%s\n%s\n", user.Password, model.HashPassword(input.Password))
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

func LogoutHandler(c *fiber.Ctx) error {
	sessionToken := c.Cookies("session_token")
	model.DB.Model(&model.User{}).Where("session_token = ?", sessionToken).Update("session_token", nil)

	c.ClearCookie("session_token")
	return c.JSON(fiber.Map{"message": "Logged out"})
}

func MeHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(model.User)
	return c.JSON(fiber.Map{
		"id":       user.ID,
		"username": user.Username,
	})
}
