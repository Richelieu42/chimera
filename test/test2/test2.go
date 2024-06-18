package main

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// CustomColorEncoderConfig returns a zapcore.EncoderConfig with colors
func CustomColorEncoderConfig() zapcore.EncoderConfig {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder // Adds color to level names
	//config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	return config
}

func main() {
	// Use the custom encoder config
	encoderConfig := CustomColorEncoderConfig()

	// Create a console encoder with the custom config
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// Create a core that writes logs to stdout
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	// Create a logger with the core
	logger := zap.New(core)

	// Example log messages
	logger.Debug("This is a debug message", zap.String("key", "value"), zap.Error(redis.Nil))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(redis.Nil))

	// Sync to flush any buffered log entries
	_ = logger.Sync()
}
