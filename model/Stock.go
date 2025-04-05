package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
)

var validStocks = []string{"AAPL", "TSLA", "GOOG", "AMZN"}

type StockQuote struct {
	CurrentPrice float64 `json:"c"`
}

type Stock struct {
	gorm.Model
	Symbol string
	Value  float64 `json:"value"`
	Amount float64 `json:"amount"`
}

// todo: refresh all values per hour
func (s *Stock) GetValue() error {
	resp, err := http.Get(fmt.Sprintf("https://finnhub.io/api/v1/quote?symbol=%s&token=%s", s.Symbol, os.Getenv("FINNHUB_API_KEY")))
	if err != nil {
		fmt.Printf("Failed to make GET request: %v\n", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Status code not ok: %d\n", resp.StatusCode)
		return errors.New("API failure")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
		return err
	}

	var quote StockQuote
	err = json.Unmarshal(body, &quote)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return err
	}

	s.Value = quote.CurrentPrice
	return nil
}

func (s *Stock) Buy(user User, dollarAmount float64) {
	user.Cash -= dollarAmount
	s.Amount += dollarAmount / s.Value
}

func (s *Stock) Sell(user User, stockAmount float64) {
	s.Amount -= stockAmount
	user.Cash += stockAmount * s.Value
}
