package console

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

var (
	L func() *zap.Logger        = zapKit.L
	S func() *zap.SugaredLogger = zapKit.S

	Sync func() = zapKit.Sync

	Debug   func(msg string, fields ...zap.Field)          = zapKit.Debug
	Debugf  func(template string, args ...interface{})     = zapKit.Debugf
	Debugw  func(msg string, keysAndValues ...interface{}) = zapKit.Debugw
	Debugln func(args ...interface{})                      = zapKit.Debugln

	// Info
	/*
	   @param fields 输出循序与传参顺序一致（并不会按字母排序）
	*/
	Info   func(msg string, fields ...zap.Field)          = zapKit.Info
	Infof  func(template string, args ...interface{})     = zapKit.Infof
	Infow  func(msg string, keysAndValues ...interface{}) = zapKit.Infow
	Infoln func(args ...interface{})                      = zapKit.Infoln

	Warn   func(msg string, fields ...zap.Field)          = zapKit.Warn
	Warnf  func(template string, args ...interface{})     = zapKit.Warnf
	Warnw  func(msg string, keysAndValues ...interface{}) = zapKit.Warnw
	Warnln func(args ...interface{})                      = zapKit.Warnln

	Error   func(msg string, fields ...zap.Field)          = zapKit.Error
	Errorf  func(template string, args ...interface{})     = zapKit.Errorf
	Errorw  func(msg string, keysAndValues ...interface{}) = zapKit.Errorw
	Errorln func(args ...interface{})                      = zapKit.Errorln

	Panic   func(msg string, fields ...zap.Field)          = zapKit.Panic
	Panicf  func(template string, args ...interface{})     = zapKit.Panicf
	Panicw  func(msg string, keysAndValues ...interface{}) = zapKit.Panicw
	Panicln func(args ...interface{})                      = zapKit.Panicln

	DPanic   func(msg string, fields ...zap.Field)          = zapKit.DPanic
	DPanicf  func(template string, args ...interface{})     = zapKit.DPanicf
	DPanicw  func(msg string, keysAndValues ...interface{}) = zapKit.DPanicw
	DPanicln func(args ...interface{})                      = zapKit.DPanicln

	Fatal   func(msg string, fields ...zap.Field)          = zapKit.Fatal
	Fatalf  func(template string, args ...interface{})     = zapKit.Fatalf
	Fatalw  func(msg string, keysAndValues ...interface{}) = zapKit.Fatalw
	Fatalln func(args ...interface{})                      = zapKit.Fatalln
)
