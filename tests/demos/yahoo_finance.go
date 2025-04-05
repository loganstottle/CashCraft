package demos

import (
	"github.com/piquette/finance-go/quote"
	"fmt"
)

func Yahoo() {
	q, err := quote.Get("AAPL")
	if err != nil {
	  panic(err)
	}

	fmt.Println(q)
}
