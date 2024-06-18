package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	prefixLogger := logger.With(zap.String("prefix", "[PREFIX]"))

	prefixLogger.Info("This is a log message")
}

func NewLogger() *zap.Logger {
	jsonFlag := true

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	var encoder zapcore.Encoder
	if jsonFlag {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// Create a core that writes logs to stdout
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	// Create a logger with the core
	logger := zap.New(core, zap.AddCaller())

	return logger
}
