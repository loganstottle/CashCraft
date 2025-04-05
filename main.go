package main

import (
	"CashCraft/controller"
)

func main() {
	controller.LoadEnv()
	controller.StartServer()
}
