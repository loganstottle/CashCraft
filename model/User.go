package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           int     `json:"userid"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Cash         float64 `json:"cash"`
	Stocks       []Stock `json:"stocks"`
	SessionToken string  `json:"session_token"`
}

func (u *User) ValuateStocks() float64 {
	var value float64 = 0

	for _, stock := range u.Stocks {
		s := StockPrice{}
		if err := DB.First(&s, "symbol = ?", stock.Symbol).Error; err != nil {
			fmt.Printf("Failed to valuate stocks for user %s: %s\n", u.Username, err)
			continue
		}

		fmt.Println(s.Value, stock.Amount)

		value += s.Value * stock.Amount
	}

	return value
}

func (u *User) HasStock(symbol string) bool {
	for _, stock := range u.Stocks {
		if stock.Symbol == symbol {
			return true
		}
	}

	return false
}
