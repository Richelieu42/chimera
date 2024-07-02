package main

import (
	"github.com/richelieu-yang/chimera/v3/src/appKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"time"
)

type defaultFatalHook struct {
}

func (h *defaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	appKit.Exit(1)
}

func main() {
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
