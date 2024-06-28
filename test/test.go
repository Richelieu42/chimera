package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 自定义的 Fatal Hook
type myFatalHook struct{}

func (h *myFatalHook) OnWrite(*zapcore.CheckedEntry, []zap.Field) {
	// 在这里实现你的钩子逻辑，比如执行清理工作
	fmt.Println("Fatal hook triggered!")
}

func main() {
	// 创建一个带有自定义 Fatal Hook 的 logger
	config := zap.NewProductionConfig()
	logger, err := config.Build(zap.WithFatalHook(&myFatalHook{}))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// 记录一些日志
	logger.Info("This is an info message")

	// 触发 fatal hook
	logger.Fatal("This is a fatal message")
}
