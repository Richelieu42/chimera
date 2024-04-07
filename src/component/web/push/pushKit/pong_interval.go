package pushKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"time"
)

var pongInterval time.Duration = time.Second * 15

func setPongInterval(interval time.Duration) error {
	if interval <= 0 {
		return errorKit.Newf("interval(%s) must be greater than 0", interval.String())
	}
	if interval < time.Millisecond*500 {
		return errorKit.Newf("interval(%s) is too small", interval.String())
	}

	pongInterval = interval
	return nil
}

func GetPongInterval() time.Duration {
	return pongInterval
}
