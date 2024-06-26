package zapKit

import (
	"go.uber.org/zap/zapcore"
	"io"
)

// NewWriteSyncer io.Writer => zapcore.WriteSyncer
func NewWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(w)
}

// NewWriteSyncerWithLock io.Writer => （线程安全的）zapcore.WriteSyncer
func NewWriteSyncerWithLock(w io.Writer) zapcore.WriteSyncer {
	ws := zapcore.AddSync(w)
	return zapcore.Lock(ws)
}
