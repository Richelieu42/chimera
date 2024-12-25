package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	L    func() *zap.Logger        = zapKit.L
	S    func() *zap.SugaredLogger = zapKit.S
	Sync func()                    = zapKit.Sync

	// SetDefaultLevel PS: 默认日志级别为 DEBUG .
	SetDefaultLevel func(level zapcore.Level) = zapKit.SetDefaultLevel
)

var (
	Debug   func(msg string, fields ...zap.Field)          = zapKit.Debug
	Debugw  func(msg string, keysAndValues ...interface{}) = zapKit.Debugw
	Debugln func(args ...interface{})                      = zapKit.Debugln

	// Info
	/*
	   @param fields 输出循序与 传参fields 顺序一致（并不会按字母排序）
	*/
	Info   func(msg string, fields ...zap.Field)          = zapKit.Info
	Infow  func(msg string, keysAndValues ...interface{}) = zapKit.Infow
	Infoln func(args ...interface{})                      = zapKit.Infoln

	Warn   func(msg string, fields ...zap.Field)          = zapKit.Warn
	Warnw  func(msg string, keysAndValues ...interface{}) = zapKit.Warnw
	Warnln func(args ...interface{})                      = zapKit.Warnln

	Error   func(msg string, fields ...zap.Field)          = zapKit.Error
	Errorw  func(msg string, keysAndValues ...interface{}) = zapKit.Errorw
	Errorln func(args ...interface{})                      = zapKit.Errorln

	Panic   func(msg string, fields ...zap.Field)          = zapKit.Panic
	Panicw  func(msg string, keysAndValues ...interface{}) = zapKit.Panicw
	Panicln func(args ...interface{})                      = zapKit.Panicln

	DPanic   func(msg string, fields ...zap.Field)          = zapKit.DPanic
	DPanicw  func(msg string, keysAndValues ...interface{}) = zapKit.DPanicw
	DPanicln func(args ...interface{})                      = zapKit.DPanicln

	Fatal   func(msg string, fields ...zap.Field)          = zapKit.Fatal
	Fatalw  func(msg string, keysAndValues ...interface{}) = zapKit.Fatalw
	Fatalln func(args ...interface{})                      = zapKit.Fatalln
)
