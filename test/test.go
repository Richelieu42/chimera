package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 自定义的钩子函数
func myHook(entry zapcore.Entry) error {
	// 在这里实现你的钩子逻辑，比如添加上下文信息
	if entry.Level == zapcore.ErrorLevel {
		// 例如：在记录错误级别的日志时添加一些额外的处理
		// 这里可以执行一些特定的逻辑
	}
	return nil
}

func main() {
	// 创建一个带有自定义钩子的logger
	config := zap.NewProductionConfig()
	logger, err := config.Build(zap.Hooks(myHook))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// 记录一些日志
	logger.Info("This is an info message")
	logger.Error("This is an error message")
}
