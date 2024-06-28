package zapKit

import (
	"go.uber.org/zap/zapcore"
	"io"
)

// NewWriteSyncer io.Writer => zapcore.WriteSyncer
/*
PS: os.File 结构体实现了 zapcore.WriteSyncer 接口.
*/
func NewWriteSyncer(w io.Writer) zapcore.WriteSyncer {
	return zapcore.AddSync(w)
}

// NewWriteSyncerWithLock io.Writer => （线程安全的）zapcore.WriteSyncer
/*
PS: os.File 结构体实现了 zapcore.WriteSyncer 接口.
*/
func NewWriteSyncerWithLock(w io.Writer) zapcore.WriteSyncer {
	ws := zapcore.AddSync(w)
	return zapcore.Lock(ws)
}

// MultiWriteSyncer 类似于 io.MultiWriter.
func MultiWriteSyncer(ws ...zapcore.WriteSyncer) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(ws...)
}
