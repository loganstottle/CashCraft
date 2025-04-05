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

	app.Static("/", "./views/home")
	app.Static("/login", "./views/login")
	app.Static("/register", "./views/register")
	SetupAuthRoutes(app)
	app.Static("/*", "./views/404")
	app.Listen(":3000")
}
