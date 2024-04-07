package pushKit

import (
	"time"
)

var (
	wsPongInterval time.Duration = time.Second * 15

	ssePongInterval time.Duration = time.Second * 15
)

func setSsePongInterval(interval time.Duration) {
	ssePongInterval = interval
}

func GetSsePongInterval() time.Duration {
	return ssePongInterval
}

func setWsPongInterval(interval time.Duration) {
	wsPongInterval = interval
}

func GetWsPongInterval() time.Duration {
	return wsPongInterval
}
