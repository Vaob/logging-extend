package datahandler //datahandler is autonomic microservice loading and updating candles every 60 seconds
// at this stage of project datahandler working with only 5-min candles, but amount of files can be expanded

import (
	"exmo-trading/app/data"
	"exmo-trading/app/query"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Handler struct {
	Symbol            string        `yaml:"Symbol"`
	Resolution        int64         `yaml:"Resolution"`
	CandlesFile       string        `yaml:"CandlesFile"`
	CandlesVolume     int64         `yaml:"CandlesVolume"`
	DataHandlerTimeout time.Duration `yaml:"DataHandlerTimeout"`
}

func (h *Handler) Nothing() {}

func (h *Handler) LoadCandles(from, to string) error {
	stringResolution := fmt.Sprintf("%d", h.Resolution)
	q := query.GetQuery{Method: "candles_history?symbol=" + h.Symbol +
		"&resolution=" + stringResolution + "&from=" + from + "&to=" + to}

	resp, err := query.Exec(&q)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	candles := &data.Candles{}
	candles.Array = make([]data.Candle, 0, h.CandlesVolume)

	err = data.ParseJson(candles, []byte(body))
	if err != nil {
		return err
	}

	if len(candles.Array