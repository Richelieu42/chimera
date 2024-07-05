package console

import (
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newLogger(skip int) *zap.Logger {
	/*
		Richelieu: 此处不直接调用 zapKit ，以防import cycle.
	*/
	encConfig := zap.NewProductionEncoderConfig()
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(encConfig)
	core := zapcore.NewCore(enc, ioKit.LockedWriteSyncerStdout, zapcore.DebugLevel)

	return zap.New(core,
		zap.WithCaller(true),
		zap.AddCallerSkip(skip),
		zap.ErrorOutput(ioKit.LockedWriteSyncerStderr),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)
}
