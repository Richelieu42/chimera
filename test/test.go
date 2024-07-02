package main

import (
	"github.com/richelieu-yang/chimera/v3/src/appKit"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type defaultFatalHook struct {
}

func (h *defaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	appKit.Exit(1)
}

func main() {
	l := zapKit.NewLogger(nil, zapKit.WithFatalHook(&defaultFatalHook{}))
	//l := zapKit.NewLogger(nil)

	l.Fatal("111", zap.String("key", "value"))
}
