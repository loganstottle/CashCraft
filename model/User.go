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

func (u *User) ValuateStocks() float64 {
	var value float64 = 0

	var stocks []Stock
	DB.Where("owner_id = ?", u.ID).Find(&stocks)
	for _, stock := range stocks {
		sp := StockPrice{}
		DB.First(&sp, "symbol = ?", stock.Symbol)
		value += stock.Amount * sp.Value
	}

	return value
}

func (u *User) GetStock(symbol string) Stock {
	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ? and symbol = ?", u.ID, symbol).Error; err != nil {
		stock.Symbol = symbol
		stock.Amount = 0
		stock.NetEarned = 0
		stock.OwnerID = u.ID
		return stock
	}

	return stock
}

func (u *User) Buy(stockSymbol string, dollars float64) error {
	if u.Cash < dollars {
		return errors.New("Player is too broke")
	}

	if dollars <= 0 {
		return errors.New("Purchase less than or equivalent to $0")
	}

	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to buy unknown stock.\n")
		return err
	}

	u.Cash -= dollars

	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ? AND symbol = ?", u.ID, stockSymbol).Error; err != nil {
		stock = Stock{
			Symbol:    stockSymbol,
			Amount:    dollars / sp.Value,
			OwnerID:   u.ID,
			NetEarned: 0,
		}
	} else {
		stock.Amount += dollars / sp.Value
	}

	stock.NetEarned -= dollars

	DB.Save(&stock)
	DB.Save(u)

	return nil
}

func (u *User) Sell(stockSymbol string, stockAmount float64) error {
	if stockAmount < 0 {
		return errors.New("Sale less than 0 stocks")
	}

	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", stockSymbol).Error; err != nil {
		fmt.Printf("Trying to sell unknown stock.\n")
		return err
	}

	stock := Stock{}
	if err := DB.First(&stock, "owner_id = ? AND symbol = ?", u.ID, stockSymbol).Error; err != nil {
		fmt.Printf("Trying to sell unowned stock\n")
		return err
	}

	// !!! Does not work !!!
	// if float64(int(stock.Amount * 1000)) / 1000.0 == stockAmount {
	// 	u.Cash += stockAmount * sp.Value
	// 	stock.Amount = 0
	// 	DB.Save(&stock)
	// 	DB.Save(u)
	// 	return nil
	// }

	stockAmount = min(stockAmount, stock.Amount)

	stock.NetEarned += stockAmount * sp.Value
	u.Cash += stockAmount * sp.Value
	stock.Amount -= stockAmount

	// if stock.Amount == 0 {
	// 	stock.NetEarned = 0
	// }

	DB.Save(&stock)
	DB.Save(u)

	return nil
}

func (u *User) Profit(symbol string) float64 {
	stock := u.GetStock(symbol)

	if stock.NetEarned == 0 && stock.Amount > 0 {
		stock.NetEarned = -stock.Amount
		DB.Save(&stock)
	}

	sp := StockPrice{}
	if err := DB.First(&sp, "symbol = ?", symbol).Error; err != nil {
		fmt.Printf("Could not get sock price %s: %s\n", symbol, err)
		return 0
	}

	return stock.Amount*sp.Value + stock.NetEarned
}
