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

/*
level < WARN，输出到: 文件日志（JSON格式）
level >= WARN，输出到: 文件日志（JSON格式） + 控制台（人类可读格式）
*/
func TestNewLogger2(t *testing.T) {
	var core1, core2 zapcore.Core

	/* 输出1: 文件 */
	{
		enc := NewEncoder(WithEncoderOutputFormatJson())

		f, err := os.Create("_test.log")
		if err != nil {
			panic(err)
		}
		defer f.Close()
		ws := zapcore.Lock(f)

		core1 = NewCore(enc, ws, zapcore.DebugLevel, zap.String("source", "0"))
	}
	/* 输出2: 控制台 */
	{

		enc := NewEncoder(WithEncoderOutputFormatConsole())
		ws := zapcore.Lock(os.Stdout)
		core2 = NewCore(enc, ws, zapcore.WarnLevel, zap.String("source", "1"))
	}

	logger := NewLogger(MultiCore(core1, core2))
	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}

type defPanicHook struct {
}

func (h *defPanicHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	//fmt.Println(entry.Message)
	//fmt.Println(entry.Level)

	entry.Write(fields...)
	panic(entry.Message)
}

func TestNewLogger3(t *testing.T) {
	//l := NewLogger(nil, WithPanicHook(&defPanicHook{}))
	l := NewLogger(nil)

	l.Panic("111", zap.String("key", "value"))
}
