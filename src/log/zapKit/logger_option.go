package zapKit

import (
	"go.uber.org/zap"
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
		// Development 是否是开发环境？会影响 zap.Logger 的DPanic方法
		/*
			true:	开发环境
			false:	（默认）生产环境
		*/
		Development bool

		// OutputType 输出类型
		OutputType outputType

		// CoreMaker 自定义Core，适用情况: n个输出（n > 2）
		/*
			PS: 如果配置了此项，第1输出（LevelEnabler、WriteSyncer）和第2输出（OtherLevelEnabler、OtherWriteSyncer）将无效.
		*/
		CoreMaker func(encoder zapcore.Encoder) zapcore.Core

		// LevelEnabler （第1输出）日志级别，支持的类型: zapcore.Level、zapcore.LevelEnabler
		LevelEnabler zapcore.LevelEnabler
		WriteSyncer  zapcore.WriteSyncer

		// OtherLevelEnabler （第2输出）想要生效，OtherLevelEnabler 和 OtherWriteSyncer必须都非nil
		OtherLevelEnabler zapcore.LevelEnabler
		// OtherWriteSyncer （第2输出）想要生效，OtherLevelEnabler 和 OtherWriteSyncer必须都非nil
		OtherWriteSyncer zapcore.WriteSyncer

		InitialFields []zap.Field

		// Caller true: 输出带有caller字段
		Caller     bool
		CallerSkip int

		MessagePrefix string

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
		Development: false,
		OutputType:  OutputTypeConsole,

		LevelEnabler: zapcore.DebugLevel,
		WriteSyncer:  zapcore.AddSync(os.Stdout),

		OtherLevelEnabler: nil,
		OtherWriteSyncer:  nil,

		Caller:        true,
		CallerSkip:    0,
		MessagePrefix: "",
		InitialFields: nil,
		EncodeLevel:   nil,
		EncodeTime:    nil,
	}

	for _, option := range options {
		option(opts)
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

func WithDevelopment(flag bool) LoggerOption {
	return func(opts *loggerOptions) {
		opts.Development = flag
	}
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

//func WithLevel(level zapcore.Level) LoggerOption {
//	return func(opts *loggerOptions) {
//		opts.Level = level
//	}
//}

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

func WithInitialFields(fields ...zap.Field) LoggerOption {
	return func(opts *loggerOptions) {
		opts.InitialFields = fields
	}
}

func WithMessagePrefix(prefix string) LoggerOption {
	return func(opts *loggerOptions) {
		opts.MessagePrefix = prefix
	}
}

// WithLevelEnabler 第 1 个输出（不能为nil）.
func WithLevelEnabler(levelEnabler zapcore.LevelEnabler) LoggerOption {
	return func(opts *loggerOptions) {
		if levelEnabler == nil {
			return
		}

		opts.LevelEnabler = levelEnabler
	}
}

// WithWriteSyncer 第 1 个输出（不能为nil）.
func WithWriteSyncer(writeSyncer zapcore.WriteSyncer) LoggerOption {
	return func(opts *loggerOptions) {
		if writeSyncer == nil {
			return
		}

		opts.WriteSyncer = writeSyncer
	}
}

// WithWriter 第 1 个输出（不能为nil）.
func WithWriter(w io.Writer) LoggerOption {
	return func(opts *loggerOptions) {
		if w == nil {
			return
		}

		opts.WriteSyncer = zapcore.AddSync(w)
	}
}

// WithOtherLevelEnabler 第 2 个输出（可以为nil）.
func WithOtherLevelEnabler(levelEnabler zapcore.LevelEnabler) LoggerOption {
	return func(opts *loggerOptions) {
		opts.OtherLevelEnabler = levelEnabler
	}
}

// WithOtherWriteSyncer 第 2 个输出（可以为nil）.
func WithOtherWriteSyncer(writeSyncer zapcore.WriteSyncer) LoggerOption {
	return func(opts *loggerOptions) {
		opts.OtherWriteSyncer = writeSyncer
	}
}
