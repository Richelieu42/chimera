package brotliKit

import (
	"bytes"
	"github.com/andybalholm/brotli"
	"io"
)

var (
	NewWriter func(dst io.Writer) *brotli.Writer = brotli.NewWriter

	NewReader func(src io.Reader) *brotli.Reader = brotli.NewReader
)

// Compress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Compress(data []byte) (compressed []byte, err error) {
	buf := bytes.NewBuffer(nil)

	brWriter := brotli.NewWriter(buf)
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

// Uncompress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Uncompress(compressed []byte) (data []byte, err error) {
	brReader := brotli.NewReader(bytes.NewBuffer(compressed))
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, brReader); err != nil {
		return
	}
	data = buf.Bytes()
	return
}
