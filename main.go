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

	TestFinance(os.Getenv("POLYGON_API_KEY"))
}
