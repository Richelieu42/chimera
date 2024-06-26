package main

import (
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	zapcore.NewMultiWriteSyncer()
	zapcore.NewLazyWith()

	defer zapKit.Sync()

	//zapKit.Info("hello", "world")
	zapKit.Info("hello", zap.String("key", "value"), zap.Bool("flag", true))
	//zapKit.Infof("hello %s", "world")
	//zapKit.Infow("hello")
	//zapKit.Infoln("hello", "world")
	//
	//zapKit.Infow("This is an info message with structured data",
	//	"key1", "value1",
	//	"key2", 42,
	//)

	logger := zapKit.NewLogger()
	logger.Info("hello", zap.String("key", "value"), zap.Bool("flag", true))

	logger.Info()

	logger.Sugar().Info("hello", zap.String("key", "value"), zap.Bool("flag", true))
}
