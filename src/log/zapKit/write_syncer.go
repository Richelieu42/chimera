package zapKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"go.uber.org/zap/zapcore"
	"io"
)

var (
	LockedWriteSyncerStdout = ioKit.LockedWriteSyncerStdout

	LockedWriteSyncerStderr = ioKit.LockedWriteSyncerStderr
)

var (
	// NewWriteSyncer io.Writer => （并发不安全的）zapcore.WriteSyncer
	/*
	   PS:
	   (1) os.File 结构体实现了 zapcore.WriteSyncer 接口;
	   (2) zapcore.WriteSyncer 接口是 io.Writer 接口的子类.
	*/
	NewWriteSyncer func(w io.Writer) zapcore.WriteSyncer = ioKit.NewWriteSyncer

	// NewLockedWriteSyncer io.Writer => （并发安全的）zapcore.WriteSyncer
	/*
	   PS:
	   (1) os.File 结构体实现了 zapcore.WriteSyncer 接口;
	   (2) zapcore.WriteSyncer 接口是 io.Writer 接口的子类.
	*/
	NewLockedWriteSyncer func(w io.Writer) zapcore.WriteSyncer = ioKit.NewLockedWriteSyncer

	// MultiWriteSyncer 类似于 io.MultiWriter.
	MultiWriteSyncer func(ws ...zapcore.WriteSyncer) zapcore.WriteSyncer = ioKit.MultiWriteSyncer
)
