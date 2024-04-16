package brotliKit

import "bytes"

type (
	lz4Options struct {
		// level 压缩级别
		level int

		// compressThreshold 压缩阈值
		/*
			单位: byte
			<= 0: ALL压缩
		*/
		compressThreshold int
	}

	Lz4Option func(opts *lz4Options)
)

func loadOptions(options ...Lz4Option) *lz4Options {
	opts := &lz4Options{
		level:             LevelDefaultCompression,
		compressThreshold: -1, /* 都压缩 */
	}

	for _, option := range options {
		option(opts)
	}

	return opts
}

func WithLevel(level int) Lz4Option {
	return func(opts *lz4Options) {
		opts.level = level
	}
}

// WithCompressThreshold 设置: 最小压缩长度阈值
func WithCompressThreshold(compressThreshold int) Lz4Option {
	return func(opts *lz4Options) {
		opts.compressThreshold = compressThreshold
	}
}

func (opts *lz4Options) Compress(data []byte) ([]byte, error) {
	if len(data) < opts.compressThreshold {
		// 不压缩
		return data, nil
	}

	buf := bytes.NewBuffer(nil)
	brWriter := NewWriterWithLevel(buf, opts.level)
	if _, err := brWriter.Write(data); err != nil {
		return nil, err
	}
	if err := brWriter.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
