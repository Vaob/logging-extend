
package strategies

import (
	"exmo-trading/app/data"
	"exmo-trading/app/trader/signals"
	"exmo-trading/app/trader/traderutils"
)

type RSI struct { // rsi strategy gives long or short signals when rsi index goes lower than 30 or higher than 70
	CandlesFile       string
	CandlesFileVolume int
	Period            int
}

func (rsi *RSI) Set(candlesFile string, candlesFileVolume int) {
	rsi.Period = 14
	rsi.CandlesFile = candlesFile
	rsi.CandlesFileVolume = candlesFileVolume
}
