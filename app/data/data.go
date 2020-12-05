package data // this package define and descibes behaviour of main data types in app

import (
	"encoding/csv"  // all data stored in csv cache ("cache/file.csv")
	"encoding/json" // need for parsing candles and responses
	"fmt"           // printing errors in console
	"os"            // working with cache files(reading, writing, rewriting