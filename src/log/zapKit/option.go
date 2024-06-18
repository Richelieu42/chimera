package zapKit

import "go.uber.org/zap/zapcore"

const (
	// OutputTypeConsole 人类可读的多行输出
	OutputTypeConsole outputType = iota

	// OutputTypeJson JSON格式输出
	OutputTypeJson
)

type (
	outputType uint8

	loggerOptions struct {
		OutputType outputType

		Level zapcore.Level

		ColorWhenConsoleEncoder bool

		// Caller true: 输出带有caller字段
		Caller bool
	}

	LoggerOption func(opts *loggerOptions)
)

func (opts *loggerOptions) IsOutputTypeConsole() bool {
	return opts.OutputType == OutputTypeConsole
}

func loadOptions(options ...LoggerOption) *loggerOptions {
	opts := &loggerOptions{
		OutputType:              OutputTypeConsole,
		Level:                   zapcore.DebugLevel,
		ColorWhenConsoleEncoder: true,
		Caller:                  true,
	}

	for _, option := range options {
		option(opts)
	}

	switch opts.OutputType {
	case OutputTypeConsole, OutputTypeJson:
	default:
		opts.OutputType = OutputTypeConsole
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

func WithColorWhenConsoleEncoder(colorWhenConsoleEncoder bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.ColorWhenConsoleEncoder = colorWhenConsoleEncoder
	}
}

func WithCaller(caller bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Caller = caller
	}
}
