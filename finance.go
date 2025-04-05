package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type StockQuote struct {
	Open      float64 `json:"o"`
	High      float64 `json:"h"`
	Low       float64 `json:"l"`
	Current   float64 `json:"c"`
	PrevClose float64 `json:"pc"`
}

func LogCurrentPrice(apiKey string, symbol string) {
	resp, err := http.Get(fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey))
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Received non-OK status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	var quote StockQuote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatalf("Error parsing JSON: %v", err)
	}

	fmt.Printf("Current price of %s stock:\n", symbol)
	fmt.Printf("Open: $%.2f\n", quote.Open)
	fmt.Printf("High: $%.2f\n", quote.High)
	fmt.Printf("Low: $%.2f\n", quote.Low)
	fmt.Printf("Current: $%.2f\n", quote.Current)
	fmt.Printf("Previous Close: $%.2f\n", quote.PrevClose)
}
