package ioKit

import "io"

// ReadReader 读取io.Reader的内容.
var ReadReader func(r io.Reader) ([]byte, error) = io.ReadAll

// ReadReaderToString 读取io.Reader的内容.
func ReadReaderToString(reader io.Reader) (string, error) {
	data, err := ReadReader(reader)
	return string(data), err
}
