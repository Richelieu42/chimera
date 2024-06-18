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
