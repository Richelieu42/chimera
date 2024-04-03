package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"github.com/sirupsen/logrus"
	"time"
)

var pool *ants.Pool

func MustSetUp(antPool *ants.Pool, logger *logrus.Logger, pongInterval time.Duration) {
	if err := Setup(antPool, logger, pongInterval); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// Setup
/*
@param antPool	需要自行决定: cap大小、是否自定义输出...
@param logger 	可以为nil
*/
func Setup(antPool *ants.Pool, logger *logrus.Logger, pongInterval time.Duration) error {
	/* pool */
	if antPool.IsClosed() {
		return errorKit.Newf("pool has already been closed")
	}
	capacity := antPool.Cap()
	if capacity > 0 {
		tag := "gte=2000"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrapf(err, "capacity(%d) of pool is invalid(tag: %s) when it's greater than zero", capacity, tag)
		}
	}
	pool = antPool

	/* logger */
	if logger != nil {
		if err := SetDefaultLogger(logger); err != nil {
			return err
		}
	}

	/* pongInterval */
	if err := setPongInterval(pongInterval); err != nil {
		return err
	}

	return nil
}

func isAvailable() error {
	if pool == nil {
		return NotSetupError
	}
	return nil
}
