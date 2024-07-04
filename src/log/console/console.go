package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

var (
	Sync func() = zapKit.Sync

	PrintBasicDetails func(loggers ...*zap.SugaredLogger) = zapKit.PrintBasicDetails

	Debug func(msg string, fields ...zap.Field) = zapKit.Debug

	Info func(msg string, fields ...zap.Field) = zapKit.Info

	Warn func(msg string, fields ...zap.Field) = zapKit.Warn

	Error func(msg string, fields ...zap.Field) = zapKit.Error

	Panic func(msg string, fields ...zap.Field) = zapKit.Panic

	DPanic func(msg string, fields ...zap.Field) = zapKit.DPanic

	Fatal func(msg string, fields ...zap.Field) = zapKit.Fatal

	Debugf func(template string, args ...interface{}) = zapKit.Debugf

	Debugw func(msg string, keysAndValues ...interface{}) = zapKit.Debugw

	Debugln func(args ...interface{}) = zapKit.Debugln

	Infof func(template string, args ...interface{}) = zapKit.Infof

	Infow func(msg string, keysAndValues ...interface{}) = zapKit.Infow

	Infoln func(args ...interface{}) = zapKit.Infoln

	Warnf func(template string, args ...interface{}) = zapKit.Warnf

	Warnw func(msg string, keysAndValues ...interface{}) = zapKit.Warnw

	Warnln func(args ...interface{}) = zapKit.Warnln

	Errorf func(template string, args ...interface{}) = zapKit.Errorf

	Errorw func(msg string, keysAndValues ...interface{}) = zapKit.Errorw

	Errorln func(args ...interface{}) = zapKit.Errorln

	DPanicf func(template string, args ...interface{}) = zapKit.DPanicf

	DPanicw func(msg string, keysAndValues ...interface{}) = zapKit.DPanicw

	DPanicln func(args ...interface{}) = zapKit.DPanicln

	Panicf func(template string, args ...interface{}) = zapKit.Panicf

	Panicw func(msg string, keysAndValues ...interface{}) = zapKit.Panicw

	Panicln func(args ...interface{}) = zapKit.Panicln

	Fatalf func(template string, args ...interface{}) = zapKit.Fatalf

	Fatalw func(msg string, keysAndValues ...interface{}) = zapKit.Fatalw

	Fatalln func(args ...interface{}) = zapKit.Fatalln
)
