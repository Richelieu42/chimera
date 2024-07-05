package ioKit

import (
	"go.uber.org/zap/zapcore"
	"io"
)

// NewLockedWriter
/*
PS:
(1) zapcore.WriteSyncer 接口继承了 io.Writer 接口.
(2) 内部做了处理，可以多次调用（虽然这么干不推荐也没意义）.

@return 支持并发 Write
*/
func NewLockedWriter(w io.Writer) zapcore.WriteSyncer {
	return zapcore.Lock(zapcore.AddSync(w))
}
