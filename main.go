package main

import (
	"CashCraft/controller"
	"CashCraft/model"
	"fmt"
)

func main() {
	controller.LoadEnv()

	for _, stock := range model.SetupStocks() {
		fmt.Println("sjdjfsdnjsdfjnsdfnjfksd")
		fmt.Println(stock)
	}

	controller.StartServer()
}
