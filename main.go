package main

import (
	"fmt"
)

func main() {
	LoadEnv()

	price := GetCurrentPrice("GOOG")
	fmt.Println(price)

	RunServer()
}
