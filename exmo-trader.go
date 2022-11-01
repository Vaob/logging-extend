package main

import (
	"exmo-trading/app/datahandler"
	"exmo-trading/app/trader"
	"exmo-trading/app/trader/strategies"
	config "exmo-trading/configs"
	"fmt"
	"path/filepath"
	"time"
)

type App struct {
	Traders      []trader.Trader
	DataHandlers []da