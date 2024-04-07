package pushKit

import (
	"time"
)

var (
	wsPongInterval time.Duration = time.Second * 15

	ssePongInterval time.Duration = time.Second * 15
)

func setWsPongInterval(interval time.Duration) {
	wsPongInterval = interval
}

func setSsePongInterval(interval time.Duration) {
	ssePongInterval = interval
}

func GetWsPongInterval() time.Duration {
	return wsPongInterval
}

func GetSsePongInterval() time.Duration {
	return ssePongInterval
}
