package main

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/request/reqKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

type defaultFatalHook struct {
}

func (h *defaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	zapKit.Exit(1)
}

func main() {

	reqKit.WithLogger(logrus.StandardLogger())
	reqKit.WithLogger(zapKit.NewSugarLogger(nil))

	zap.IncreaseLevel()
	zapcore.NewIncreaseLevelCore()

	zap.S()

	go func() {
		zapKit.Info("[goroutine] sleep starts")
		time.Sleep(time.Second * 3)
		zapKit.Info("[goroutine] sleep ends")
	}()

	zapKit.Info("[main] sleep starts")
	time.Sleep(time.Second)
	zapKit.Info("[main] sleep ends")
	runtime.Goexit()
}
