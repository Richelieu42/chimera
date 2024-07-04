package console

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	l *zap.Logger
	s *zap.SugaredLogger
)

func init() {
	/*
		Richelieu: 此处不直接调用 zapKit ，以防import cycle.
	*/
	encConfig := zap.NewProductionEncoderConfig()
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	enc := zapcore.NewConsoleEncoder(encConfig)
	ws := zapcore.Lock(os.Stdout)
	core := zapcore.NewCore(enc, ws, zapcore.DebugLevel)

	l = zap.New(core, zap.WithCaller(true), zap.AddCallerSkip(1))
	s = l.Sugar()
}

//// ReplaceGlobalLoggers
///*
//PS: 有需要的话，应该在应用初始化时调用此方法，即在最前面设置全局logger.
//*/
//func ReplaceGlobalLoggers(logger *zap.Logger) {
//	if logger == nil {
//		return
//	}
//
//	l = logger
//	s = logger.Sugar()
//}

func Sync() {
	_ = l.Sync()
	_ = s.Sync()
}

func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	l.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	l.DPanic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	l.Fatal(msg, fields...)
}
