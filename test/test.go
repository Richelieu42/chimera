package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// JSON Encoder
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// Console Encoder
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	// 创建Logger，使用JSON格式
	loggerJSON := zap.New(zapcore.NewCore(jsonEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel))
	loggerJSON.Info("This is a JSON formatted log", zap.String("key", "value"))

	// 创建Logger，使用Console格式
	loggerConsole := zap.New(zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel))
	loggerConsole.Info("This is a Console formatted log", zap.String("key", "value"))
}
