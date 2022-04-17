package strategies

import (
	"exmo-trading/app/data"
	"exmo-trading/app/trader/signals"
	"exmo-trading/app/trader/traderutils"
)

type BollingerBands struct { // rsi strategy gives long or short signals when rsi index goes lower than 30 or higher than 70
	CandlesFile       string
	CandlesFileVolume int
	Period            int
	Factor			  int
}

fun