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
var validStocksNames = []string{"Apple", "Tesla", "Google", "Amazon"}

type StockQuote struct {
	CurrentPrice float64 `json:"c"`
}

type StockPrice struct {
	Symbol string `json:"symbol"`
	Name   string
	Value  float64 `json:"value"`
}

type Stock struct {
	gorm.Model
	Symbol string
	Amount float64 `json:"amount"`
	UserID int     // foreign key GORM requirement
	User   User    // foreign key GORM requirement
}

func SetupStocks() []StockPrice {
	var stocks []StockPrice

	for i, stockSymbol := range validStocks {
		s := StockPrice{stockSymbol, validStocksNames[i], 0}
		s.UpdatePrice()
		stocks = append(stocks, s)
		if err := DB.First(&s, "symbol = ?", stockSymbol).Error; err != nil {
			DB.Create(&s)
		} else {
			stock := DB.Model(StockPrice{}).Where("symbol = ?", stockSymbol)
			stock.Update("value", s.Value)
			DB.Save(&s)
		}
	}

	u := User{}
	DB.First(&u, "username = ?", "test")

	s := Stock{}
	s.Symbol = "AAPL"
	s.Amount = 0

	s.Buy(u, 1000)

	fmt.Println(u.ValuateStocks())

	return stocks
}

// todo: refresh all values per hour
func (s *StockPrice) UpdatePrice() error {
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

func (s *Stock) Buy(user User, dollarAmount float64) error {
	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", s.Symbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	user.Cash -= dollarAmount
	s.Amount += dollarAmount / sp.Value

	if !user.HasStock(s.Symbol) {
		s.UserID = user.ID
		user.Stocks = append(user.Stocks, *s)
		DB.Save(&user)
		DB.Create(&s)
	} else {
		DB.Update("stocks", &s)
	}

	return nil
}

func (s *Stock) Sell(user User, stockAmount float64) {
	s.Amount -= stockAmount
	user.Cash += stockAmount * s.Amount
}
