package ioKit

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
)

func NewReader(s []byte) io.Reader {
	return bytes.NewReader(s)
}

func NewReaderFromString(str string) io.Reader {
	return strings.NewReader(str)
}

// NewFileReader
/*
！！！：要在外部手动调用 *os.File 的Close方法.

@return os.File 结构体 实现了 io.Reader 接口
*/
func NewFileReader(filePath string) (*os.File, error) {
	return os.Open(filePath)
}

// ToBufioReader io.Reader 接口 => *bufio.Reader
/*
作用: 可以调用 bufio.Reader 结构体的方法（因为 io.Reader 接口就一个Read方法）.
*/
func ToBufioReader(reader io.Reader) *bufio.Reader {
	return bufio.NewReader(reader)
}

// ReadFromReader 读取io.Reader的内容（io.Reader => []byte）
func ReadFromReader(reader io.Reader) ([]byte, error) {
	return io.ReadAll(reader)
}

// ReadStringFromReader 读取io.Reader的内容（io.Reader => string）
func ReadStringFromReader(reader io.Reader) (string, error) {
	data, err := io.ReadAll(reader)
	return string(data), err
}

func Read() {

}
