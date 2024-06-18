package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	logger, _ := zap.NewDevelopment()

	logger.Debug("This is a debug message.")
	logger.Info("This is an info message.")
	logger.Warn("This is an warn message.")
	logger.Error("This is an error message.")

	zap.LevelFlag()
	zap.NewAtomicLevel()
	zap.ParseAtomicLevel()

	zapcore.ParseLevel()
	zapcore.LevelOf()

	zapcore.NewConsoleEncoder()
	zapcore.NewJSONEncoder()

	zap.NewDevelopmentEncoderConfig()

	zapcore.NewJSONEncoder()
	zapcore.NewJSONEncoder()
}
