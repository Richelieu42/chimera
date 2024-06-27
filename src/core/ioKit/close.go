package ioKit

import (
	"io"
	"os"
)

// TryToClose 尝试关闭（如果实现了io.Closer接口的话）.
/*
@param writers 	(1) 可以为nil（即不传参）
				(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func TryToClose(objs ...any) (err error) {
	for _, obj := range objs {
		if closer, ok := obj.(io.Closer); ok {
			tmpErr := closeSingle(closer)
			if tmpErr != nil && err == nil {
				err = tmpErr
			}
		}
	}
	return err
}

// Close
/*
PS: 就算循环过程中返回了非nil的error，也要继续向下循环（尽可能多地关闭）.

@param closers (1) 可以为nil（即不传参）；(2) 其中可以有nil
@return 发生error的话，返回第一个
*/
func Close(closers ...io.Closer) (err error) {
	for _, closer := range closers {
		tmpErr := closeSingle(closer)
		if tmpErr != nil && err == nil {
			err = tmpErr
		}
	}
	return err
}

func closeSingle(closer io.Closer) error {
	if closer != nil {
		switch closer {
		case os.Stdin, os.Stdout, os.Stderr:
			// 这3种不关闭
			return nil
		default:
			return closer.Close()
		}
	}
	return nil
}
