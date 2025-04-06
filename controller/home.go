package controller

import (
	"CashCraft/model"

	"github.com/gofiber/fiber/v2"
	//"strconv"
	"fmt"
)

func FormatBalance(amount float64) string {
	var result string
	balance_str := fmt.Sprintf("%.2f", (amount - 0.005)) // This looks weird, but prevents rounding up
	// This doesn't steal money, it just shows you have your amount, or your amount but smaller by one cent

	if amount < 1000 {
		return "$" + balance_str // If you don't have one thousand dollar, commas are not needed
	}

	for i := len(balance_str) - 1; i >= 0; i-- {

		if i < len(balance_str)-4 && (len(balance_str)-i-1)%3 == 0 { // Commaizing code
			result = "," + result
		}

		result = string(balance_str[i]) + result
	}

	return "$" + result // dollar sign added before the number after commas are added
}

func SetupHomeRoutes(app *fiber.App) {
	app.Get("/", GetHome) // Totally could have put this in the controller.go file, but this makes the code look nicer
}

func GetHome(c *fiber.Ctx) error {
	var user model.User

	// If you try to go to the home page without being logged in, it takes you to the login screen (which has a button to go register)
	if err := model.DB.First(&user, "username = ? and session_token = ?", c.Cookies("username"), c.Cookies("session_token")).Error; err != nil {
		return c.Redirect("/login")
	}

	var stocksData string
	for _, stock := range model.SetupStocks() {
		stocksData += fmt.Sprintf("<strong>%s (%s)</strong>: %s<br>", stock.Name, stock.Symbol, FormatBalance(stock.Value))
	} // This turns the stock data into a string - because we had issues with go templates

	return c.Render("./view/home/index.html", fiber.Map{ // Data that we feed into the home page
		"Username":       user.Username,
		"Balance":        FormatBalance(user.Cash),
		// "StockValuation": user.ValuateStocks(),
		"StocksData":     stocksData,
	})
}
