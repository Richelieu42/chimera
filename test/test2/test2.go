package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

// 自定义编码器来给日志消息添加前缀
type prefixEncoder struct {
	zapcore.Encoder
	prefix string
}

func (pe *prefixEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 给msg字段加上前缀
	entry.Message = pe.prefix + entry.Message
	return pe.Encoder.EncodeEntry(entry, fields)
}

func main() {
	// 创建一个文件日志输出
	file, _ := os.Create("logfile.log")
	fileWriter := zapcore.AddSync(file)

	// 创建一个控制台输出
	consoleWriter := zapcore.AddSync(os.Stdout)

	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 创建控制台编码器
	consoleEncoder := zapcore.NewConsoleEncoder(encoderConfig)

	// 创建带前缀的控制台编码器
	prefix := "myPrefix: "
	prefixConsoleEncoder := &prefixEncoder{
		Encoder: consoleEncoder,
		prefix:  prefix,
	}

	// 创建正常日志级别的核心
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})
	infoCore := zapcore.NewCore(prefixConsoleEncoder, fileWriter, infoLevel)

	// 创建错误日志级别的核心
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	errorCore := zapcore.NewCore(prefixConsoleEncoder, zapcore.NewMultiWriteSyncer(fileWriter, consoleWriter), errorLevel)

	// 合并核心
	core := zapcore.NewTee(infoCore, errorCore)

	// 创建Logger
	logger := zap.New(core, zap.AddCaller())

	defer logger.Sync() // 刷新所有缓冲

	// 测试日志
	logger.Info("This is an info message")
	logger.Error("This is an error message")
}
