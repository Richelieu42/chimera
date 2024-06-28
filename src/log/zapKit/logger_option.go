package zapKit

type (
	loggerOptions struct {
		// Development 是否是开发环境？会影响 zap.Logger 的DPanic方法
		/*
			true:	开发环境
			false:	（默认）生产环境
		*/
		Development bool

		// Caller true: 输出带有caller字段
		Caller     bool
		CallerSkip int
	}

	LoggerOption func(opts *loggerOptions)
)

func loadOptions(options ...LoggerOption) *loggerOptions {
	opts := &loggerOptions{
		Development: false,
		Caller:      true,
		CallerSkip:  0,
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
