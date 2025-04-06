package controller

import (
	"CashCraft/model"
	"html/template"

	"github.com/gofiber/fiber/v2"
	//"strconv"
	"fmt"
)

func FormatBalance(amount float64) string {
	var result string
	balance_str := fmt.Sprintf("%.2f", amount)
	// <<<<<<< HEAD

	// =======
	// >>>>>>> 35e13ebcd09545d598d5fd7ec6f4b76e7ccae19d
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

	// This code is used to calculate the net worth of the user
	cashBalance := user.Cash
	netWorth := cashBalance + user.ValuateStocks()

	var myStockData string
	var stocksData string
	for _, stock := range model.GetStocks() {
		myStockData += fmt.Sprintf("<strong>%s (%s)</strong>: %s (%s) <button id=\"buy-%s\" class=\"buy\">Buy</button> <button id=\"sell-%s\" class=\"sell\">Sell</button><br>", stock.Name, stock.Symbol, FormatBalance(user.GetStock(stock.Symbol))[1:], FormatBalance(stock.Value*user.GetStock(stock.Symbol)), stock.Symbol, stock.Symbol)
		stocksData += fmt.Sprintf("<strong>%s (%s)</strong>: %s<br>", stock.Name, stock.Symbol, FormatBalance(stock.Value))
	} // This turns the stock data into a string - because we had issues with go templates

	return c.Render("home/index", fiber.Map{ // Data that we feed into the home page
		"Username":    user.Username,
		"NetWorth":    FormatBalance(netWorth),
		"CashBalance": FormatBalance(cashBalance),
		"StocksData":  template.HTML(stocksData),
		"MyStocks":    template.HTML(myStockData),
	})
}
