package zapKit

import (
	"go.uber.org/zap/zapcore"
	"io"
)

// NewWriteSyncer
func NewWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(w)
}

// NewWriteSyncerWithLock
func NewWriteSyncerWithLock(w io.Writer) zapcore.WriteSyncer {
	ws := zapcore.AddSync(w)
	return zapcore.Lock(ws)
}
