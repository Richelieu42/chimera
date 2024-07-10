package main

import (
	"go.uber.org/zap"
	"log"
)

func main() {
	// 配置 Zap 日志库
	cfg := zap.NewProductionConfig()
	logger, _ := cfg.Build()

	// 将标准库 log 输出重定向到 Zap 日志库
	zap.RedirectStdLog(logger)

	// 使用标准库 log 打印日志
	log.Println("This is a log message using standard log package")

	// 使用 Zap 日志库打印日志
	logger.Info("This is a log message using Zap logger")

	// 关闭日志库
	defer logger.Sync()
}
