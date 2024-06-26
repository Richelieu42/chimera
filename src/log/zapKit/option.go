package zapKit

import (
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

const (
	// OutputTypeConsole 人类可读的多行输出
	OutputTypeConsole outputType = iota

	// OutputTypeJson JSON格式输出
	OutputTypeJson
)

type (
	outputType uint8

	loggerOptions struct {
		WriteSyncer zapcore.WriteSyncer

		// OutputType 输出类型
		OutputType outputType

		// Level 日志级别
		Level zapcore.Level

		// Caller true: 输出带有caller字段
		Caller bool

		// Development 是否是开发环境？会影响 zap.Logger 的DPanic方法
		/*
			true:	开发环境
			false:	生产环境
		*/
		Development bool

		EncodeLevel zapcore.LevelEncoder
		EncodeTime  zapcore.TimeEncoder
	}

	LoggerOption func(opts *loggerOptions)
)

func (opts *loggerOptions) IsOutputTypeConsole() bool {
	return opts.OutputType == OutputTypeConsole
}

func loadOptions(options ...LoggerOption) *loggerOptions {
	opts := &loggerOptions{
		WriteSyncer: nil,
		OutputType:  OutputTypeConsole,
		Level:       zapcore.DebugLevel,
		Caller:      true,
		Development: false,
		EncodeLevel: nil,
		EncodeTime:  nil,
	}

	for _, option := range options {
		option(opts)
	}

	// WriteSyncer
	if opts.WriteSyncer == nil {
		opts.WriteSyncer = zapcore.AddSync(os.Stdout)
	}
	// outputType
	switch opts.OutputType {
	case OutputTypeConsole, OutputTypeJson:
	default:
		opts.OutputType = OutputTypeConsole
	}
	// EncodeLevel
	if opts.EncodeLevel == nil {
		if opts.IsOutputTypeConsole() {
			opts.EncodeLevel = zapcore.CapitalColorLevelEncoder
		} else {
			opts.EncodeLevel = zapcore.CapitalLevelEncoder
		}
	}
	// EncodeTime
	if opts.EncodeTime == nil {
		if opts.IsOutputTypeConsole() {
			opts.EncodeTime = zapcore.ISO8601TimeEncoder
		} else {
			opts.EncodeTime = zapcore.EpochTimeEncoder
		}
	}

	return opts
}

func WithOutputTypeJson() LoggerOption {
	return func(opts *loggerOptions) {
		opts.OutputType = OutputTypeJson
	}
}

func WithOutputTypeConsole() LoggerOption {
	return func(opts *loggerOptions) {
		opts.OutputType = OutputTypeConsole
	}
}

func WithLevel(level zapcore.Level) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Level = level
	}
}

func WithCaller(caller bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Caller = caller
	}
}

func WithEncodeLevel(encodeLevel zapcore.LevelEncoder) LoggerOption {
	return func(opts *loggerOptions) {
		opts.EncodeLevel = encodeLevel
	}
}

func WithEncodeTime(encodeTime zapcore.TimeEncoder) LoggerOption {
	return func(opts *loggerOptions) {
		opts.EncodeTime = encodeTime
	}
}

func WithWriteSyncer(writeSyncer zapcore.WriteSyncer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.WriteSyncer = writeSyncer
	}
}

// WithWriter 设置输出
func WithWriter(w io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		if w != nil {
			opts.WriteSyncer = NewWriteSyncer(w)
		}
	}
}

// WithLockedWriter 设置输出（会给输出加锁，并发安全地!!!）
func WithLockedWriter(w io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		if w != nil {
			opts.WriteSyncer = NewWriteSyncerWithLock(w)
		}
	}
}

func WithDevelopment(flag bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Development = flag
	}
}
