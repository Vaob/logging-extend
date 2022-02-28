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
	DataHandlerTimeout time.Duratio