package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
	//"strconv"
	"fmt"
)

func SetupHomeRoutes(app *fiber.App) {
	app.Get("/", GetHome)
}

func GetHome(c *fiber.Ctx) error {
	var user model.User
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err != nil {
		return c.Redirect("/login")
	}

//	bal := fmt.Sprintf("$%s", strconv.FormatInt(int64(user.Cash), 10))
	bal := fmt.Sprintf("$%.2f", user.Cash)

	return c.Render("./view/home/index.html", fiber.Map{
		"Balance": bal,
	})
}
