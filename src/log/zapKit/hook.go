package zapKit

import (
	"go.uber.org/zap/zapcore"
)

// defaultFatalHook 默认的 fatal hook
type defaultFatalHook struct {
}

func (h *defaultFatalHook) OnWrite(entry *zapcore.CheckedEntry, fields []zapcore.Field) {
	Exit(1)
}
