package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/robfig/cron/v3"
)

// To eventually be changed to allow like - the S&P 500 - or any one on the NYSE stock exchange

var ValidStocks = []string{"AAPL", "TSLA", "GOOG", "AMZN", "NVDA", "MSFT", "META", "COST", "DIS", "NFLX", "PLTR", "WMT", "V", "MA", "KO", "MCD", "PEP", "ADBE", "LMT"}                                                                      // Stocks accepted for everything
var ValidStocksNames = []string{"Apple", "Tesla", "Google", "Amazon", "Nvidia", "Microsoft", "Meta", "Costco", "Disney", "Netflix", "Palantir", "Walmart", "Visa", "Mastercard", "Coca-Cola", "McDonald's", "PepsiCo", "Adobe", "Lockheed"} // The names of those stocks

type StockQuote struct { // a struct that holds the current price for a stock
	Change        float64 `json:"d"`
	ChangePercent float64 `json:"dp"`
	CurrentPrice  float64 `json:"c"`
}

type StockPrice struct { // a struct that holds the stocks symbol, name, and cost
	Symbol             string `json:"symbol"`
	Name               string
	Value              float64 `json:"value"`
	DailyChange        float64 `json:"daily_change"`
	DailyChangePercent float64 `json:"daily_change_percent"`
	ID                 uint    `gorm:"primary_key" json:"id"`
}

type Stock struct { // We have owner id to tie who owns each one
	gorm.Model
	Symbol  string
	Amount  float64 `json:"amount"`
	OwnerID uint
}

var MarketState bool

func (s *StockPrice) UpdatePrice() error { // API call with lots of error checking to update price of the stock passed into it
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
	s.DailyChange = quote.Change
	s.DailyChangePercent = quote.ChangePercent
	return nil
}

func (sp *StockPrice) Up() bool {
	return sp.DailyChange > 0
}

func (sp *StockPrice) GenerateStatusString() string {
	var emoji string

	if MarketState == false {
		emoji = ""
	} else if sp.Up() {
		emoji = "📈"
	} else {
		emoji = "📉"
	}

	return fmt.Sprintf("%s $%.2f %.2f%%", emoji, math.Abs(sp.DailyChange), math.Abs(sp.DailyChangePercent))
}

func SetupStocks() {
	if MarketState == false {
		return
	}
	for i, stockSymbol := range ValidStocks {
		var s StockPrice
		if err := DB.First(&s, "symbol = ?", stockSymbol).Error; err != nil {
			s.Symbol = stockSymbol
			s.Name = ValidStocksNames[i]
			DB.Create(&s)
		}

		s.UpdatePrice()
		DB.Save(&s)
	}

	fmt.Println("stocks updated")
}

func OpenMarket(c *cron.Cron, cronj string) {
	c.AddFunc(cronj, func() {
		MarketState = true
		fmt.Println("The market is now OPEN")
	})
}

func CloseMarket(c *cron.Cron, cronj string) {
	c.AddFunc(cronj, func() {
		MarketState = false
		fmt.Println("The market is now CLOSED")
	})
}

func SetupStocksCron() {
	MarketState = true
	SetupStocks() // always grab stocks when market is open, immediately recycle (conditional below)
	if time.Now().Hour() > 0 && time.Now().Hour() < 8 {
		MarketState = false
		fmt.Println("The market is now CLOSED")
	} else {
		MarketState = true
		fmt.Println("The market is now OPEN")
	}
	stateCron := cron.New()
	apiCron := cron.New()

	// TODO: redo crons for non-daylight savings
	OpenMarket(stateCron, "0 8 * * 1-5")  // Standard Market Opening
	CloseMarket(stateCron, "0 0 * * 1-5") // Standard Market Closing

	CloseMarket(stateCron, "* 17 3 7 *")   // July 3rd Market Close
	CloseMarket(stateCron, "* 17 28 11 *") // Black Friday Market Close
	CloseMarket(stateCron, "* 17 24 12 *") // Christmas Eve Market Close

	CloseMarket(stateCron, "* 4 1,9,20 1 *") // January Holidays
	CloseMarket(stateCron, "* 4 17 2 *")     // President's Day
	CloseMarket(stateCron, "* 4 18 4 *")     // Good Friday
	CloseMarket(stateCron, "* 4 26 5 *")     // Memorial Day
	CloseMarket(stateCron, "* 4 19 6 *")     // Juneteenth
	CloseMarket(stateCron, "* 4 4 7 *")      // Independence Day
	CloseMarket(stateCron, "* 4 1 9 *")      // Labor Day
	CloseMarket(stateCron, "* 4 27 11 *")    // Thanksgiving
	CloseMarket(stateCron, "* 4 25 12 *")    // Christmas

	apiCron.AddFunc("@every 1m", func() {
		SetupStocks()
	})

	stateCron.Start()
	apiCron.Start()
}

func GetStocks() []StockPrice {
	var stocks []StockPrice
	DB.Find(&stocks)
	return stocks
}
