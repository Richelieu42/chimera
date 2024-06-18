package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	url := "https:///www.baidu.com"

	zap.New()

	zap.Config{}
	zap.SamplingConfig{}

	zap.NewDevelopment()
	zap.NewProduction()
	zap.NewExample()
	zap.NewNop()

	zap.L()

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)
}
