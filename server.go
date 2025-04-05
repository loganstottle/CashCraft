package main

import (
	"github.com/gofiber/fiber/v2"
)

func RunServer() {
	app := fiber.New()

	app.Static("/", "./views/home")
	app.Static("/login", "./views/login")
	app.Static("/*", "./views/404")

	app.Listen(":3000")
}
