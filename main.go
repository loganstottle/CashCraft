package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env")
	}

	price := GetCurrentPrice(os.Getenv("FINNHUB_API_KEY"), "GOOG")
	fmt.Println(price)
}
