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

func (u *User) Buy(stockSymbol string, dollars float64) error {
	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	u.Cash -= dollars

	if !u.HasStock(stockSymbol) {

		return nil
	}

	// user.Cash -= dollars
	// s.Amount += dollars / sp.Value

	// if !user.HasStock(s.Symbol) {
	// 	s.UserID = user.ID
	// 	user.Stocks = append(user.Stocks, *s)
	// 	DB.Save(&user)
	// 	DB.Create(&s)
	// } else {
	// 	DB.Update("stocks", &s)
	// }

	// return nil
}
