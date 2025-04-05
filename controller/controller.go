package controller

import (
	"CashCraft/model"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Failed to load .env\n")
	}
}

func StartServer() {
	model.ConnectDatabase()

	app := fiber.New()

	app.Static("/", "./view/home")
	app.Static("/login", "./view/login")
	app.Static("/register", "./view/register")
	SetupAuthRoutes(app)
	app.Static("/*", "./view/404")
	app.Listen(":3000")
}
