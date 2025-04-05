package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Could not load .env")
	}

	LogCurrentPrice(os.Getenv("FINNHUB_API_KEY"), "AAPL")
}
