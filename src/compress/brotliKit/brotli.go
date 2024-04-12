package brotliKit

import (
	"bytes"
	"github.com/andybalholm/brotli"
	"io"
)

// Compress
/*
PS: 参考 github.com/andybalholm/brotli 中的 "example_test.go".
*/
func Compress(data []byte) (compressed []byte, err error) {
	buf := bytes.NewBuffer(nil)

	writer := brotli.NewWriter(buf)
	_, err = writer.Write(data)
	if err != nil {
		return
	}
	if err = writer.Close(); err != nil {
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
	reader := brotli.NewReader(bytes.NewBuffer(compressed))
	buf := bytes.NewBuffer(nil)
	if _, err = io.Copy(buf, reader); err != nil {
		return
	}
	data = buf.Bytes()
	return
}
