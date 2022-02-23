package datahandler //datahandler is autonomic microservice loading and updating candles every 60 seconds
// at this stage of project datahandler working with only 5-min candles, but amount of files can be expanded

import (
	"exmo-trading/app/data"
	"exmo-