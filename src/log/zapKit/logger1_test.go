package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"go.uber.org/zap"
	"testing"
)

func TestNewLogger(t *testing.T) {
	// 日志文件会生成在当前.go文件的同目录下
	writer, err := fileKit.CreateInAppendMode("aaa.log")
	if err != nil {
		panic(err)
	}

	logger, err := NewLogger1(writer, zap.InfoLevel)
	if err != nil {
		panic(err)
	}

	// 由于日志级别，下一行不会输出
	logger.Debug("Debug")
	logger.Info("Info")
	logger.Warn("Warn")
	logger.Error("Error")
}
