
package trader // trader is autonomic microservice that launch strategies and handling signals by making events

import (
	"exmo-trading/app/data"
	"exmo-trading/app/trader/signals"
	"exmo-trading/app/trader/strategies"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Event struct { // event is need to making trades from strategy signals
	//we making trade relying on previous trades and last price
	Action    string      //handled signal:long or short
	LastPrice float64     //actual price of traded pair
	Trades    data.Trades //array of open trades
}

type Trader struct {
	strategies.Strategy
	Context TraderContext
}

type TraderContext struct {
	TradingPair       string  `yaml:"TradingPair"`       //trading pair; exaple: "BTC_USDT"
	TradesFile        string  `yaml:"TradesFile"`        //trades file name; example: "/cache/trades.csv"
	TradesHistoryFile string  `yaml:"TradesHistoryFile"` //trades history file name; example: "/cache/trades-history.csv"
	CandlesFile       string  `yaml:"CandlesFile"`       //candles file name; example: "/cache/5min-candles.csv"
	CandlesFileVolume int     `yaml:"CandlesFileVolume"` // volume of candles in candles file; default: 250
	TradesFileVolume  int     `yaml:"TradesFileVolume"`  // max amount of open trades; default: 10 (means 10 trades)
	StopLimitPercent  float64 `yaml:"StopLimitPercent"`  // 0.01 means 1%, when price change goes higher
	// or lower this percent, we close unprofitable trade
	// default: 0.01
	TradeAmount   float64       `yaml:"TradeAmount"`   // trade volume; example: 100 USDT in BTC_USDT pair
	TraderTimeout time.Duration `yaml:"TraderTimeout"` // time of trader reload; default: 30
}

func (t *TraderContext) Nothing() {

}

func (e *Event) GetLastPrice(ctx *TraderContext) error {
	url := "https://api.exmo.me/v1.1/ticker"
	method := "POST"
	payload := strings.NewReader("")
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}