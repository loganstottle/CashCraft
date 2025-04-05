package model

import (
	"fmt"
	"errors"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID           int     `gorm:"primaryKey"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Cash         float64 `json:"cash"`
	Stocks       []Stock `gorm:"many2many;stocks;"`
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

func (u *User) GetStock(symbol string) (Stock, error) {
	for _, stock := range u.Stocks {
		if stock.Symbol == symbol {
			return stock, nil
		}
	}

	return Stock{}, errors.New("no stock exists")
}

func (u *User) Buy(stockSymbol string, dollars float64) error {
	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	u.Cash -= dollars

	stock, stockErr := u.GetStock(stockSymbol)
	if stockErr != nil {
		stock := Stock{}
		stock.Symbol = stockSymbol
		stock.Amount = 0
		stock.UserID = u.ID
	}

	stock.Amount += dollars / sp.Value

	u.Stocks = append(u.Stocks, stock)

	DB.Save(&u)
	if stockErr == nil {
		DB.Save(&stock)
	} else {
		DB.Create(&stock)
	}

	return nil
}
