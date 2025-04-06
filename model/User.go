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
		value += stock.Amount * sp.Value
	}

	return value, nil
}

func (u *User) GetStock(symbol string) float64 {
	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ?", u.ID).Error; err != nil {
		return 0
	}

	return stock.Amount
}

func (u *User) Buy(stockSymbol string, dollars float64) error {
	if u.Cash <= dollars {
		return errors.New("Player is too broke")
	}

	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	u.Cash -= dollars

	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ?", u.ID).Error; err != nil {
		stock := Stock{
			Symbol:  stockSymbol,
			Amount:  dollars / sp.Value,
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

func (u *User) Sell(stockSymbol string, stockAmount float64) error {
	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to sell unknown stock.\n")
		return err
	}

	u.Cash += stockAmount * sp.Value

	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ?", u.ID).Error; err != nil {
		fmt.Printf("Trying to sell unowned stock\n")
		return err
	}
	stock.Amount -= stockAmount
	DB.Save(&stock)

	DB.Save(u)

	return nil
}
