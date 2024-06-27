package zapKit

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

// TestNewLogger
/*
第1个输出: DEBUG级别及以上的输出，输出到文件
第2个输出: WARN级别及以上的输出，输出到控制台

总结: 	(1) < WARN级别，输出到: 文件
		(2) >= WARN级别，输出到: 文件、控制台
*/
func TestNewLogger(t *testing.T) {
	f, err := os.Create("_test.log")
	if err != nil {
		panic(err)
	}
	ws1 := zapcore.AddSync(f)
	ws2 := zapcore.AddSync(os.Stdout)

	logger := NewLogger(
		WithOutputTypeConsole(),
		WithMessagePrefix("[TEST] "),
		WithInitialFields(zap.Bool("c", true)),

		WithLevelEnabler(zapcore.DebugLevel),
		WithWriteSyncer(ws1),

		WithOtherLevelEnabler(zapcore.WarnLevel),
		WithOtherWriteSyncer(ws2),
	)
	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}
