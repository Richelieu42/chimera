package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func main() {
	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	// 创建控制台编码器
	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	tmp0 := zapcore.AddSync(os.Stdout)
	tmp1 := zapcore.AddSync(os.Stderr)

	// 创建正常日志级别的核心
	infoLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level < zapcore.ErrorLevel
	})
	infoCore := zapcore.NewCore(encoder, tmp0, infoLevel)

	// 创建错误日志级别的核心
	errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
		return level >= zapcore.ErrorLevel
	})
	errorCore := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(tmp0, tmp1), errorLevel)

	// 合并核心
	core := zapcore.NewTee(infoCore, errorCore)

	// 创建Logger
	logger := zap.New(core, zap.AddCaller())

	defer logger.Sync() // 刷新所有缓冲

	// 测试日志
	logger.Debug("debug")
	logger.Info("info")
	logger.Warn("warn")
	logger.Error("error")
}
