/*
This is unused code from when we where hoping to use the yahoo api - but that went
no where so it isn't included at all in docs about the project

If you found this, congratulations! This is an easter egg

Its dangerous to go alone!
Take this...
\****__              ____
  |    *****\_      --/ *\-__
  /_          (_    ./ ,/----'
      \__         (_./  /
         \__           \___----^__
        _/   _                  \
 |    _/  __/ )\"\ _____         *\
 |\__/   /    ^ ^       \____      )
  \___--"                    \_____ )

Its the cute little dragon dude!
*/

package demos

import (
	"fmt"

	"github.com/piquette/finance-go/quote"
)

func Yahoo() {
	q, err := quote.Get("AAPL")
	if err != nil {
		panic(err)
	}

	fmt.Println(q)
}
