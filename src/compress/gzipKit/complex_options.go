package gzipKit

import (
	"compress/gzip"
)

const (
	// DefaultLevel 默认的压缩级别，速度最快
	DefaultLevel = gzip.BestSpeed

	// DefaultContentLengthThreshold 默认的最小压缩长度阈值
	DefaultContentLengthThreshold = -1
)

type (
	gzipOptions struct {
		// level 压缩级别
		level int

		// contentLengthThreshold 最小压缩长度阈值
		/*
			单位: byte
			<= 0: ALL压缩
		*/
		contentLengthThreshold int
	}

	GzipOption func(opts *gzipOptions)
)

func loadOptions(options ...GzipOption) *gzipOptions {
	opts := &gzipOptions{
		level:                  DefaultLevel,
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

func (opts *gzipOptions) Compress(data []byte) ([]byte, error) {
	if err := AssertValidLevel(opts.level); err != nil {
		return nil, err
	}

	if len(data) < opts.contentLengthThreshold {
		// (1) 不进行压缩
		return data, nil
	}
	// (2) 进行压缩
	return Compress(data, opts.level)
}
