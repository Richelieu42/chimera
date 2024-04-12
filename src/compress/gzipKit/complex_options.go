package gzipKit

const (
	DefaultContentLengthThreshold = 256
)

type (
	gzipOptions struct {
		// level 压缩级别
		level int

		// contentLengthThreshold 最小压缩长度阈值，单位: byte
		contentLengthThreshold int
	}

	GzipOption func(opts *gzipOptions)
)

func loadOptions(options ...GzipOption) *gzipOptions {
	opts := &gzipOptions{
		level:                  1,
		contentLengthThreshold: DefaultContentLengthThreshold,
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithLevel(level int) GzipOption {
	return func(opts *gzipOptions) {
		opts.level = level
	}
}

// WithContentLengthThreshold 设置: 最小压缩长度阈值
func WithContentLengthThreshold(contentLengthThreshold int) GzipOption {
	return func(opts *gzipOptions) {
		opts.contentLengthThreshold = contentLengthThreshold
	}
}
