package ioKit

import "io"

// ReadAll 读取 io.Reader实例 的内容.
var ReadAll func(r io.Reader) ([]byte, error) = io.ReadAll

// ReadAllToString 读取 io.Reader实例 的内容.
func ReadAllToString(reader io.Reader) (string, error) {
	data, err := ReadAll(reader)
	return string(data), err
}
