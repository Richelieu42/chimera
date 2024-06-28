package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"go.uber.org/zap/zapcore"
)

// NewCore
/*
@param enc 			不能为nil
@param ws			不能为nil
@param levelEnabler	(1) 不能为nil
					(2) 可以是多种:
						(a) zapcore.Level 类型（级别 >= 此值的才会输出）
							e.g.
							zapcore.DebugLevel
							zapcore.InfoLevel
							zapcore.WarnLevel
							zapcore.ErrorType
							zapcore.PanicLevel
							zapcore.DPanicLevel
							zapcore.FatalLevel
							zapcore.InvalidLevel
						(b) zap.LevelEnablerFunc 类型（更加地自定义）
							e.g.
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
