package zapKit

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	{
		logger := NewLogger()
		defer logger.Sync()

		logger.Debug("This is a debug message", zap.String("key", "value"))
		logger.Info("This is an info message")
		logger.Warn("This is a warning message")
		logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
	}

	fmt.Println("------")

	{
		logger := NewLogger(WithOutputTypeJson())
		defer logger.Sync()

		logger.Debug("This is a debug message", zap.String("key", "value"))
		logger.Info("This is an info message")
		logger.Warn("This is a warning message")
		logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
	}
}
