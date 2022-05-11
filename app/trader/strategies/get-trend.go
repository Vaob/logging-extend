
package strategies

import (
	"exmo-trading/app/data"
	"exmo-trading/app/trader/signals"
	"exmo-trading/app/trader/traderutils"
)

type GetTrend struct { // get trend gives bull or bear signals when actual price lower or hihger than ma200