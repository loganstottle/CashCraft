package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

var validStocks = []string{"AAPL", "TSLA", "GOOG", "AMZN"}

type StockQuote struct {
	CurrentPrice float64 `json:"c"`
}

func GetCurrentPrice(apiKey string, symbol string) float64 {
	resp, err := http.Get(fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", symbol, apiKey))
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status code not ok: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	var quote StockQuote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	return quote.CurrentPrice
}
