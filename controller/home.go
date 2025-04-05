package controller

import (
	"github.com/gofiber/fiber/v2"
)

func SetupHomeRoutes(app *fiber.App) {
	app.Get("/", GetHome)
}

func GetHome(c *fiber.Ctx) error {
	return c.Render("./view/home/index.html", fiber.Map{})
}
