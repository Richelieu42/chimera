package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"go.uber.org/zap/zapcore"
)

// NewCore
/*
@param levelEnabler (1) zapcore.Level 类型实现了 zapcore.LevelEnabler 接口
					(2) e.g. 自定义
						// 创建错误日志级别的核心
						errorLevel := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
							return level >= zapcore.ErrorLevel
						})
*/
func NewCore(enc zapcore.Encoder, ws zapcore.WriteSyncer, levelEnabler zapcore.LevelEnabler) zapcore.Core {
	return zapcore.NewCore(enc, ws, levelEnabler)
}

func NewLazyWith(core zapcore.Core, fields []zapcore.Field) zapcore.Core {
	return zapcore.NewLazyWith(core, fields)
}

func NewIncreaseLevelCore(core zapcore.Core, levelEnabler zapcore.LevelEnabler) (zapcore.Core, error) {
	return zapcore.NewIncreaseLevelCore(core, levelEnabler)
}

// MultiCore
/*
@return != nil
*/
func MultiCore(cores ...zapcore.Core) zapcore.Core {
	cores = sliceKit.RemoveZeroValues(cores)

	return zapcore.NewTee(cores...)
}
