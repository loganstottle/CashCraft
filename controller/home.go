package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
	//"strconv"
	"fmt"
)

func FormatBalance(amount float64) string {
	var result string
	balance_str := fmt.Sprintf("%.2f", amount)

	if amount < 1000 {
		return "$" + balance_str
	}

	for i := len(balance_str) - 1; i >= 0; i-- {

		if i < len(balance_str)-4 && (len(balance_str)-i-1)%3 == 0 {
			result = "," + result
		}

		result = string(balance_str[i]) + result
	}

	return "$" + result
}

func SetupHomeRoutes(app *fiber.App) {
	app.Get("/", GetHome)
}

func GetHome(c *fiber.Ctx) error {
	var user model.User
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err != nil {
		return c.Redirect("/login")
	}

	//	bal := fmt.Sprintf("$%s", strconv.FormatInt(int64(user.Cash), 10))

	var stocksData string
	for _, stock := range model.SetupStocks() {
		stocksData += fmt.Sprintf("<strong>%s (%s)</strong>: %s<br>", stock.Name, stock.Symbol, FormatBalance(stock.Value))
	}

	return c.Render("./view/home/index.html", fiber.Map{
		"Username":   user.Username,
		"Balance":    FormatBalance(user.Cash),
		"StocksData": stocksData,
	})
}
