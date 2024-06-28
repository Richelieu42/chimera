package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapcore.RegisterHooks()

	zap.Hooks()

	zap.WithPanicHook()

	zap.WithFatalHook()

	zapcore.WriteThenPanic
	zapcore.WriteThenFatal
	zapcore.WriteThenGoexit
}
