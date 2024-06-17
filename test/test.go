package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// 创建一个新的 logrus Logger 实例
	logger := logrus.New()

	// 配置 TextFormatter 带上颜色
	logger.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,                  // 强制颜色输出
		FullTimestamp:    true,                  // 显示完整时间戳
		TimestampFormat:  "2006-01-02 15:04:05", // 自定义时间戳格式
		QuoteEmptyFields: false,                 // 对空字段加引号
	})

	// 示例日志输出
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message")
}
