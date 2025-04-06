package unit_tests

import (
	"CashCraft/controller"
	"testing"
)

// We where working on adding test, currently only have this one, but we did manual ones too

func TestFormatting(t *testing.T) { //This tests the formattinmg of text
	var input = []float64{0.0, 1.1, 5.35, 9.99, 10.0, 19.99, 99.99, 99.0, 100.0, 199.99, 1000.0, 1999.99, 10000.0, 15931.59, 100593.53, 1935135.53, 19531531.23, 51935193.93}
	var expected = []string{"$0.00", "$1.10", "$5.35", "$9.99", "$10.00", "$19.99", "$99.99", "$99.00", "$100.00", "$199.99", "$1,000.00", "$1,999.99", "$10,000.00", "$15,931.59", "$100,593.53", "$1,935,135.53", "$19,531,531.23", "$51,935,193.93"}
	for i, v := range input {
		if controller.FormatBalance(v) != expected[i] {
			t.Errorf("FormatBalance: %s != %s", controller.FormatBalance(v), expected[i])
		}
	}
}
