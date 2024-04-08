package pushKit

import (
	"github.com/panjf2000/ants/v2"
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"github.com/sirupsen/logrus"
)

var (
	setupFlag = atomicKit.NewBool(false)

	// pushPool 并发执行推送任务
	pushPool *ants.Pool
)

func MustSetUp(antPool *ants.Pool, logger *logrus.Logger, options ...Option) {
	if err := Setup(antPool, logger, options...); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

// Setup
/*
@param antPool	用来并发执行推送任务
				(1) 不能为nil
				(2) 需要自行决定: cap大小、是否自定义输出...
@param logger 	可以为nil
@param options
*/
func Setup(antPool *ants.Pool, logger *logrus.Logger, options ...Option) (err error) {
	defer func() {
		setupFlag.Store(err == nil)
	}()

	opts := loadOptions(options...)

	/* antPool */
	if err := interfaceKit.AssertNotNil(antPool, "antPool"); err != nil {
		return err
	}
	if antPool.IsClosed() {
		return errorKit.Newf("antPool has already been closed")
	}
	capacity := antPool.Cap()
	if capacity > 0 {
		tag := "gte=100"
		if err := validateKit.Var(capacity, tag); err != nil {
			return errorKit.Wrapf(err, "capacity(%d) of antPool is invalid(tag: %s) when it's greater than zero", capacity, tag)
		}
	}
	pushPool = antPool

	/* logger */
	if logger != nil {
		if err := SetDefaultLogger(logger); err != nil {
			return err
		}
	}

	/* pongInterval */
	setWsPongInterval(opts.WsPongInterval)
	setSsePongInterval(opts.SsePongInterval)

	return nil
}

func CheckSetup() error {
	if !setupFlag.Load() {
		return NotSetupError
	}
	return nil
}
