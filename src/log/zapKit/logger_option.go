package zapKit

import (
	"go.uber.org/zap/zapcore"
	"os"
)

type (
	loggerOptions struct {
		// Development 是否是开发环境？会影响 zap.Logger 的DPanic方法
		/*
			true:	开发环境
			false:	（默认）生产环境
		*/
		Development bool

		// ErrorOutput
		/*
			默认: zapcore.Lock(os.Stderr)
		*/
		ErrorOutput zapcore.WriteSyncer

		// Caller true: 输出带有caller字段
		Caller     bool
		CallerSkip int

		AddStacktrace zapcore.LevelEnabler

		Clock zapcore.Clock

		Fields []zapcore.Field
	}

	LoggerOption func(opts *loggerOptions)
)

func loadOptions(options ...LoggerOption) *loggerOptions {
	opts := &loggerOptions{
		Development:   false,
		ErrorOutput:   zapcore.Lock(os.Stderr),
		Caller:        true,
		CallerSkip:    0,
		AddStacktrace: zapcore.ErrorLevel, /* Error及以上的日志输出，会附带堆栈信息 */
		Clock:         zapcore.DefaultClock,
		Fields:        nil,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithDevelopment(flag bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Development = flag
	}
}

func WithErrorOutput(writeSyncer zapcore.WriteSyncer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.ErrorOutput = writeSyncer
	}
}

func WithCaller(caller bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Caller = caller
	}
}

func WithCallerSkip(skip int) LoggerOption {
	return func(opts *loggerOptions) {
		opts.CallerSkip = skip
	}
}

func WithAddStacktrace(levelEnabler zapcore.LevelEnabler) LoggerOption {
	return func(opts *loggerOptions) {
		opts.AddStacktrace = levelEnabler
	}
}

func WithClock(clock zapcore.Clock) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Clock = clock
	}
}

func WithFields(fields ...zapcore.Field) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Fields = fields
	}
}
