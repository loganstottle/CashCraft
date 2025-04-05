package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	ID           int     `json:"userid"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Cash         float64 `json:"cash"`
	Stocks       []Stock `json:"stocks"`
	SessionToken string  `json:"session_token"`
}

// func (u *User) ValuateStocks() float64 {
// 	var value float64 = 0

// 	for _, stock := range u.Stocks {
// 		value += [stock.Symbol] * stock.Amount
// 	}

// 	return value
// }
