package zapKit

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	enc := NewEncoder()
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}
