package controller

import (
	"CashCraft/model"
	"html/template"

	"fmt"

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

	// sort users (array of model.User) by their value of i.Cash (where i is an individual iteration of model.User looped through)
	// users[0], users[1] = users[1], users[0]
	for i := 0; i < len(users)-2; i++ {
		for j := i + 1; j < len(users)-1; j++ {
			x, _ := users[i].ValuateStocks()
			y, _ := users[j].ValuateStocks()
			// fmt.Printf("%f > %f\n", y + users[j].Cash, x + users[i].Cash)
			if y+users[j].Cash > x+users[i].Cash {
				users[i], users[j] = users[j], users[i]
			}
		}
	}

	// for i := 0; i < n - 1; i++ {
	// 	for j := 0; j < n - i - 1; j++ {
	// 		if users[j].Cash > users[j + 1].Cash {
	// 			users[j], users[j + 1] = users[j + 1], users[j]
	// 		}
	// 	}
	// }

	for i, v := range users {
		switch i {
		case 0:
			leaderboard += "<span style=\"color: goldenrod\">"
		case 1:
			leaderboard += "<span style=\"color: grey\">"
		case 2:
			leaderboard += "<span style=\"color: saddlebrown\">"
		default:
			leaderboard += "<span>"
		}
		leaderboard += fmt.Sprintf("%d. ", i+1)
		leaderboard += v.Username + "</span><span>: "
		x, _ := v.ValuateStocks()
		leaderboard += fmt.Sprintf("%s (%s)", FormatBalance(v.Cash), FormatBalance(x+v.Cash))
		leaderboard += "<br>"
	}

	return c.Render("leaderboard/index", fiber.Map{
		"leaderboard": template.HTML(leaderboard),
	})
}
