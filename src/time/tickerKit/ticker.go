package tickerKit

import "time"

var NewTicker func(d time.Duration) *time.Ticker = time.NewTicker
