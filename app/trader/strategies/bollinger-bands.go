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

func (bb *BollingerBands) Set(candlesFile string, candlesFileVolume int) {
	bb.Period = 40
	bb.Factor = 2
	bb.CandlesFile = candlesFile
	bb.CandlesFileVolume = candlesFileVolume
}

func (bb *BollingerBands) Solve(c *data.Candles, topborder, bottomborder []float64) string {
	length := len(c.Array)
	if length != 0 {
		if c.Array[length-1].Close > topborder[length-1] {
			return signals.Short
		}