package brotliKit

import (
	"bytes"
	"github.com/andybalholm/brotli"
	"io"
)

var (
	NewWriter func(dst io.Writer) *brotli.Writer = brotli.NewWriter

	NewWriterWithLevel func(dst io.Writer, level int) *brotli.Writer = brotli.NewWriterLevel

	NewReader func(src io.Reader) *brotli.Reader = brotli.NewReader
)

// Compress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Compress(data []byte, options ...Lz4Option) (compressed []byte, err error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}

// Decompress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Decompress(compressed []byte) ([]byte, error) {
	// brReader := brotli.NewReader(bytes.NewBuffer(compressed))

	brReader := NewReader(bytes.NewReader(compressed))

	/* 方法1 */
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, brReader); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil

	/* 方法2 */
	//return io.ReadAll(brReader)
}
