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
	DataHandlers []datahandler.Handler
}

func PrepareTrader(path, configName string, strategy strategies.Strategy) trader.Trader {
	trader := trader.Trader{}
	err := config.Load(&trader.Context, configName)
	if err != nil {
		fmt.Println(err)
	}
	strategy.Set(path + trader.Context.CandlesFile, trader.Context.CandlesFileVolume)
	trader.Context.TradesFile = path + trader.Context.TradesFile
	trader.Context.TradesHistoryFile = path + trader.Context.TradesHistoryFile
	trader.Context.CandlesFile = path + trader.Context.CandlesFile
	trader.Strategy = strategy
	return trader
}

func PrepareDataHandler(path, configName string) datahandler.Handler {
	datahandler := datahandler.Handler{}
	err := config.Load(&datahandler, configName)
	if err != nil {
		fmt.Println(err)
	}
	datahandler.CandlesFile = path + datahan