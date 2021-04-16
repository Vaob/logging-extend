package data // this package define and descibes behaviour of main data types in app

import (
	"encoding/csv"  // all data stored in csv cache ("cache/file.csv")
	"encoding/json" // need for parsing candles and responses
	"fmt"           // printing errors in console
	"os"            // working with cache files(reading, writing, rewriting)
	"strconv"       // need for working with csv
)

type Candles struct {
	Array []Candle `json:"candles"` // candles array
}

type Candle struct {
	Time   int64   `json:"t"` // close time of candle
	Open   float64 `json:"o"` // open price of candle
	Close  float64 `json:"c"` // close price of candle
	High   float64 `json:"h"` // high price of candle
	Low    float64 `json:"l"` // low price of candle
	Volume float64 `json:"v"` // traded volume of candle
}

type Trades struct {
	Array []Trade `json:"trades"` // trades array
}

type Trade struct {
	Id         int64   `json:"id"`          // id of trade
	Action     string  `json:"action"`      // long or short
	OpenPrice  float64 `json:"open-price"`  // starting price of trade
	ClosePrice float64 `json:"close-price"` // finished price of trade
	Quantity   float64 `json:"quantity"`    // trade volume, example : amount of usdt when you buy btc
	StopLimit  float64 `json:"stop-limit"`  // when price goes to stop limit, trade closed
	Status     string  `json:"status"`      // opened or closed
}

type Ticker struct {
	BuyPrice  string `json:"buy_price"` // ticker need for getting actual price of trading pair
	SellPrice string `json:"sell_price"`
	LastTrade string `json:"last_trade"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Avg       string `json:"avg"`
	Vol       string `json:"vol"`
	VolCurr   string `json:"vol_curr"`
	Updated   int64  `json:"updated"`
}

type Data interface {
	CanParseJson()
	Read(string) error
	Write(string) error
}

func (t *Ticker