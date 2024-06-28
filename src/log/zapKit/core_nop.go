package zapKit

import "go.uber.org/zap/zapcore"

func NewNopCore() zapcore.Core {
	return zapcore.NewNopCore()
}
