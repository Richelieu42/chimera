package main

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
)

func main() {
	logger := zapKit.NewLogger(zapKit.WithOutputTypeJson())
	defer logger.Sync()

	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))
}
