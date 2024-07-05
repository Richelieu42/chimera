package console

import (
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
	core := zapcore.NewCore(enc, LockedWriteSyncerStdout, zapcore.DebugLevel)

	return zap.New(core, zap.WithCaller(true), zap.AddCallerSkip(skip), zap.ErrorOutput(LockedWriteSyncerStderr))
}
