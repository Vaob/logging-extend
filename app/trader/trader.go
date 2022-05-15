
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