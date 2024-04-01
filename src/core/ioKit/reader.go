package ioKit

import (
	"bufio"
	"bytes"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"io"
	"os"
	"strings"
)

var (
	// NewReader
	/*
		PS: bytes.Reader 结构体实现了 io.Reader 接口.
	*/
	NewReader func(b []byte) *bytes.Reader = bytes.NewReader

	// NewReaderFromString
	/*
		PS: strings.Reader 结构体实现了 io.Reader 接口.
	*/
	NewReaderFromString func(s string) *strings.Reader = strings.NewReader
)

// NewReaderFromPath
/*
!!!: 要在外部手动调用 *os.File 的Close方法.

PS: os.File 结构体 实现了 io.Reader 接口.

@param path 文件（或目录）的路径
*/
func NewReaderFromPath(path string) (*os.File, error) {
	if err := fileKit.AssertExist(path); err != nil {
		return nil, err
	}
	return os.Open(path)
}

// NewBufioReader 带缓冲的Reader.
/*
PS: bufio.Reader 结构体 实现了 io.Reader 接口.
*/
func NewBufioReader(reader io.Reader, bufSizeArgs ...int) *bufio.Reader {
	if len(bufSizeArgs) == 0 {
		// 默认缓冲大小: 4096
		return bufio.NewReader(reader)
	}
	return bufio.NewReaderSize(reader, bufSizeArgs[0])
}
