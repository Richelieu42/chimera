package ioKit

import (
	"bytes"
)

var (
	NewBuffer func(s []byte) *bytes.Buffer = bytes.NewBuffer

	NewBufferString func(str string) *bytes.Buffer = bytes.NewBufferString

	// NewReadWriter bytes.Buffer 结构体 实现了 io.ReadWriter 接口
	/*
		@param s 可以为nil
	*/
	NewReadWriter func(s []byte) *bytes.Buffer = bytes.NewBuffer

	// NewReadWriterFromString bytes.Buffer 结构体 实现了 io.ReadWriter 接口
	/*
		@param str 可以为""
	*/
	NewReadWriterFromString func(str string) *bytes.Buffer = bytes.NewBufferString
)
