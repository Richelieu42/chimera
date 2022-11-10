package ioKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"io"
	"os"
)

type writeCloser struct {
	writer io.Writer
}

func (wc *writeCloser) Write(p []byte) (int, error) {
	if wc == nil {
		return 0, errorKit.Simple("wc == nil")
	}

	return wc.writer.Write(p)
}

func (wc *writeCloser) WriteString(s string) (int, error) {
	if wc == nil {
		return 0, errorKit.Simple("wc == nil")
	}

	return io.WriteString(wc.writer, s)
}

func (wc *writeCloser) Close() error {
	if wc == nil {
		return errorKit.Simple("wc == nil")
	}

	return CloseWriter(wc.writer)
}

// WrapToWriteCloser io.Writer => io.WriteCloser
func WrapToWriteCloser(writer io.Writer) (io.WriteCloser, error) {
	switch writer {
	case nil:
		return nil, errorKit.Simple("writer == nil")
	case os.Stdout:
		fallthrough
	case os.Stderr:
		// 这2种情况必须继续执行以封装，以免被误操作而关闭
	default:
		// 如果本来就是 io.WriteCloser类型，直接返回
		if tmp, ok := writer.(io.WriteCloser); ok {
			return tmp, nil
		}
	}

	return &writeCloser{
		writer: writer,
	}, nil
}

// multiWriteCloser （实现了 io.WriteCloser 接口）
/*
	可以实现: 同时输出到 文件 和 控制台...
*/
type multiWriteCloser struct {
	writeClosers []io.WriteCloser
}

func (mwc *multiWriteCloser) Write(p []byte) (int, error) {
	if mwc == nil {
		return 0, errorKit.Simple("mwc == nil")
	}
	dataSize := len(p)

	for _, writer := range mwc.writeClosers {
		n, err := writer.Write(p)
		if err != nil {
			return 0, err
		}
		if n != dataSize {
			return 0, io.ErrShortWrite
		}
	}
	return dataSize, nil
}

func (mwc *multiWriteCloser) WriteString(s string) (int, error) {
	if mwc == nil {
		return 0, errorKit.Simple("mwc == nil")
	}
	dataSize := len(s)

	for _, writer := range mwc.writeClosers {
		n, err := io.WriteString(writer, s)
		if err != nil {
			return 0, err
		}
		if n != dataSize {
			return 0, io.ErrShortWrite
		}
	}
	return dataSize, nil
}

func (mwc *multiWriteCloser) Close() error {
	if mwc == nil {
		return errorKit.Simple("mwc == nil")
	}

	var firstErr error

	for _, closer := range mwc.writeClosers {
		// 就算循环过程中返回了非nil的error，也要继续向下循环（关闭尽可能多的io.Closer），返回第一个非nil的error
		err := closer.Close()
		if err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

// MultiWriteCloser 将 多个io.WriteCloser 组合为 1个io.WriteCloser
/*
PS: 参考了 io.MultiWriter().

@param writeClosers (1) 可以不传参; (2) 有传参的话，其中不能有nil，否则会返回error.
*/
func MultiWriteCloser(writeClosers ...io.WriteCloser) (io.WriteCloser, error) {
	all := make([]io.WriteCloser, 0, len(writeClosers))

	for _, writeCloser := range writeClosers {
		if writeCloser == nil {
			return nil, errorKit.Simple("nil in writeClosers")
		}

		if mwc, ok := writeCloser.(*multiWriteCloser); ok {
			all = append(all, mwc.writeClosers...)
		} else {
			all = append(all, writeCloser)
		}
	}
	return &multiWriteCloser{
		writeClosers: all,
	}, nil
}
