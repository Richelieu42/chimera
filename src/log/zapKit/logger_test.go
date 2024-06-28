package zapKit

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := NewLogger(nil)
	/*
		WithAddStacktrace(zapcore.WarnLevel): Warn及以上的日志输出，会附带堆栈信息
	*/
	//l := NewLogger(nil, WithAddStacktrace(zapcore.WarnLevel))

	defer l.Sync()

	l.Debug("This is a debug message", zap.String("key", "value"))
	l.Info("This is an info message")
	l.Warn("This is a warning message")
	l.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}

/*
core 和 logger 都能添加 自定义Fields.
*/
func TestNewLogger1(t *testing.T) {
	encoder := NewEncoder()
	// 确保多个goroutine在写入日志时不会发生竞态条件
	ws := zapcore.Lock(os.Stdout)
	core0 := NewCore(encoder, ws, zapcore.DebugLevel, zap.String("source", "X"))
	core1 := NewCore(encoder, ws, zapcore.DebugLevel, zap.String("source", "Y"))
	core := MultiCore(core0, core1)
	l := NewLogger(core, WithFields(zap.String("source", "O")))

	l.Debug("This is a debug message")
	l.Info("This is an info message")
	l.Warn("This is a warning message")
	l.Error("This is an error message0\nThis is an error message1")
}

//// TestNewLogger
///*
//第1个输出: DEBUG级别及以上的输出，输出到文件
//第2个输出: WARN级别及以上的输出，输出到控制台
//
//总结: 	(1) < WARN级别，输出到: 文件
//		(2) >= WARN级别，输出到: 文件、控制台
//*/
//func TestNewLogger(t *testing.T) {
//	f, err := os.Create("_test.log")
//	if err != nil {
//		panic(err)
//	}
//	ws1 := zapcore.AddSync(f)
//	ws2 := zapcore.AddSync(os.Stdout)
//
//	logger := NewLogger(
//		WithOutputTypeConsole(),
//		WithMessagePrefix("[TEST] "),
//		WithInitialFields(zap.Bool("c", true)),
//
//		WithLevelEnabler(zapcore.DebugLevel),
//		WithWriteSyncer(ws1),
//
//		WithOtherLevelEnabler(zapcore.WarnLevel),
//		WithOtherWriteSyncer(ws2),
//	)
//	logger.Debug("This is a debug message", zap.String("key", "value"))
//	logger.Info("This is an info message")
//	logger.Warn("This is a warning message")
//	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
//}
