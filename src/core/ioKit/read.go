package ioKit

import "io"

// ReadAll 读取 io.Reader实例 的内容（从 io.Reader实例 中读取所有可用的数据直到 EOF（文件结束符），并返回一个包含所有读取数据的字节切片）.
/*
!!!: 由于 io.ReadAll 将整个文件内容加载到内存中，因此对于大文件或资源有限的情况，请谨慎使用，以免造成"内存溢出".在这种情况下，应考虑采用逐块读取数据的流式处理方式.
*/
var ReadAll func(r io.Reader) ([]byte, error) = io.ReadAll

func ReadAllToString(reader io.Reader) (string, error) {
	data, err := ReadAll(reader)
	return string(data), err
}

// ReadFull 尝试从 io.Reader 中读取指定长度的数据，并确保至少读取这么多数据，否则它会返回一个错误.
/*
PS: 这个函数通常用于需要固定长度数据块的协议或格式中，例如在处理网络包或解码结构化的二进制数据时.
*/
var ReadFull func(reader io.Reader, buf []byte) (n int, err error) = io.ReadFull

// ReadAtLeast 试图从 io.Reader 中读取至少指定数量的字节，如果实际读取的数量少于最小要求值，则返回错误.
/*
PS: 不同于 ReadFull 要求读取确切长度的数据，ReadAtLeast 只保证最少读到多少字节，所以即使读取到了超过 min 字节的数据，也可能会立即返回.
*/
var ReadAtLeast func(reader io.Reader, buf []byte, min int) (n int, err error) = io.ReadAtLeast
