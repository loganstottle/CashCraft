package controller

import (
	"CashCraft/model"
	"html/template"

	"fmt"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func SetupLeaderboardRoutes(app *fiber.App) {
	app.Get("/leaderboard", GetLeaderboard)
}

func GetLeaderboard(c *fiber.Ctx) error {
	var user model.User
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err != nil {
		return c.Redirect("/login")
	}

	var leaderboard string
	var users []model.User
	model.DB.Find(&users)

	users = users[:10]

	// sort users (array of model.User) by their value of i.Cash (where i is an individual iteration of model.User looped through)
	sort.Slice(users, func(i, j int) bool {
		return (users[j].Cash + users[j].ValuateStocks()) < (users[i].Cash + users[i].ValuateStocks())
	})

	for i, v := range users {
		leaderboard += fmt.Sprintf("<span class=\"number\">%d.</span> ", i+1)
		switch i {
		case 0:
			leaderboard += "<span class=\"name\" style=\"color: goldenrod\">"
		case 1:
			leaderboard += "<span class=\"name\" style=\"color: grey\">"
		case 2:
			leaderboard += "<span class=\"name\" style=\"color: saddlebrown\">"
		default:
			leaderboard += "<span class=\"name\" style=\"number\">"
		}
		leaderboard += v.Username + "</span> - "
		leaderboard += fmt.Sprintf("<span class=\"stock\">%s</span>", FormatBalance(v.Cash+v.ValuateStocks()))
		leaderboard += "<br>"
	}

	return c.Render("leaderboard/index", fiber.Map{
		"leaderboard": template.HTML(leaderboard),
	})
}
