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
	Volume float64 `json:"v"