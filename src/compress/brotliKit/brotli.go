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

func Compress(data []byte) (compressed []byte, err error) {
	// 默认压缩级别: 6
	return CompressWithLevel(data, LevelDefaultCompression)
}

// CompressWithLevel
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func CompressWithLevel(data []byte, level int) (compressed []byte, err error) {
	buf := bytes.NewBuffer(nil)

	brWriter := NewWriterWithLevel(buf, level)
	_, err = brWriter.Write(data)
	if err != nil {
		return
	}
	if err = brWriter.Close(); err != nil {
		return
	}
	compressed = buf.Bytes()
	return
}

// Decompress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Decompress(compressed []byte) (data []byte, err error) {
	//brReader := brotli.NewReader(bytes.NewBuffer(compressed))

	brReader := brotli.NewReader(bytes.NewReader(compressed))
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, brReader); err != nil {
		return
	}
	data = buf.Bytes()
	return
}
