package controller

import (
	"CashCraft/model"
	"html/template"
	"math"

	"github.com/gofiber/fiber/v2"
	//"strconv"
	"fmt"
)

func FormatBalance(amount float64) string {
	var result string
	balance_str := fmt.Sprintf("%.2f", amount)
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

	cashBalance := user.Cash
	netWorth := cashBalance + user.ValuateStocks()

	var myStockData string
	var stocksData string

	//todo: templates (different engine may have fixed it)
	for _, stockPrice := range model.GetStocks() {
		myStock := user.GetStock(stockPrice.Symbol)

		myStockData += fmt.Sprintf("<div class=\"stock-info\"><strong>%s (%s)</strong> <span style=\"color: #666\">-</span> <span style=\"text-decoration: underline\">%.3f</span> shares ", stockPrice.Name, stockPrice.Symbol, myStock.Amount)

		if myStock.Amount > 0 {
			if math.Abs(user.Profit(stockPrice.Symbol)) < 0.01 {
				myStockData += fmt.Sprintf("<span style=\"color: #666\">(%s)</span>", FormatBalance(stockPrice.Value*myStock.Amount))
			} else if user.Profit(stockPrice.Symbol) > 0 {
				myStockData += fmt.Sprintf("<span style=\"color: #666\">(%s</span> <span style=\"color: #2e2\">+%s profit<span style=\"color: #666\">)</span>", FormatBalance(stockPrice.Value*myStock.Amount), FormatBalance(user.Profit(stockPrice.Symbol)))
			} else {
				myStockData += fmt.Sprintf("<span style=\"color: #666\">(%s</span> <span style=\"color: #f22\">-%s lost<span style=\"color: #666\">)</span>", FormatBalance(stockPrice.Value*myStock.Amount), FormatBalance(user.Profit(stockPrice.Symbol)))
			}
		}

		myStockData += fmt.Sprintf(" <div class=\"btns-container\"><button id=\"buy-%s\" class=\"buy\">Buy</button> <button id=\"sell-%s\" class=\"sell\">Sell</button></div><br></div>", stockPrice.Symbol, stockPrice.Symbol)
		stocksData += fmt.Sprintf("<strong>%s (%s)</strong> <span style=\"color: #666\">-</span> ", stockPrice.Name, stockPrice.Symbol)

		if stockPrice.Up() {
			stocksData += fmt.Sprintf("<span style=\"color: #2e2; font-weight: bold;\">%s</span> ", FormatBalance(stockPrice.Value))
		} else {
			stocksData += fmt.Sprintf("<span style=\"color: #f22; font-weight: bold;\">%s</span> ", FormatBalance(stockPrice.Value))
		}

		stocksData += fmt.Sprintf("<span style=\"color: #666\">(%s)</span><br>", stockPrice.GenerateStatusString())
	} // This turns the stock data into a string - because we had issues with go templates

	return c.Render("home/index", fiber.Map{ // Data that we feed into the home page
		"Username":    user.Username,
		"NetWorth":    FormatBalance(netWorth),
		"CashBalance": FormatBalance(cashBalance),
		"StocksData":  template.HTML(stocksData),
		"MyStocks":    template.HTML(myStockData),
	})
}
