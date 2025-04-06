package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	Cash         float64 `json:"cash"`
	SessionToken string  `json:"session_token"`
}

func (u *User) ValuateStocks() (float64, error) {
	var value float64 = 0

	var stocks []Stock
	DB.Where("owner_id = ?", u.ID).Find(&stocks)
	for _, stock := range stocks {
		sp := StockPrice{}
		if err := DB.First(&sp, "symbol = ?", stock.Symbol).Error; err != nil {
			return -1, errors.New("Trying to evaluate nonexistent stock")
		}
		fmt.Println(stock.Amount * sp.Value)
		value += stock.Amount * sp.Value
	}

	return value, nil
}

// func (u *User) GetStock(symbol string) (Stock, error) {
	// for _, stock := range u.Stocks {
	// 	if stock.Symbol == symbol {
	// 		return stock, nil
	// 	}
	// }

	// return Stock{}, errors.New("no stock exists")
// }

// func (u *User) SetStock(stock Stock) error {
	// for i, s := range u.Stocks {
	// 	if s.Symbol == stock.Symbol {
	// 		u.Stocks[i].Amount = stock.Amount
	// 		fmt.Println(u.Stocks[i].Amount)
	// 		return nil
	// 	}
	// }

	// return errors.New("no stock exists")
// }

func (u *User) Buy(stockSymbol string, dollars float64) error {
	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	u.Cash -= dollars

	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ?", u.ID).Error; err != nil {
		stock := Stock{
			Symbol: stockSymbol,
			Amount: dollars / sp.Value,
			OwnerID: u.ID,
		}

		DB.Create(&stock)
	} else {
		stock.Amount += dollars / sp.Value
		DB.Save(&stock)
	}

	DB.Save(u)

	return nil
}
