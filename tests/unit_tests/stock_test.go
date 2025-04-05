package unit_tests

import (
	"CashCraft/model"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.

type StockPrice = model.StockPrice

func TestSetupStocks(t *testing.T) {
	var msg []StockPrice = model.SetupStocks()
	t.Errorf(msg[0].Symbol)
}
