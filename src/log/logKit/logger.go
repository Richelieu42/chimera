package logKit

import (
	"github.com/richelieu-yang/chimera/v2/src/file/fileKit"
	"io"
	"log"
	"os"
)

// NewLogger
/*
@param flag e.g. os.O_CREATE|os.O_WRONLY|os.O_APPEND
*/
var NewLogger func(out io.Writer, prefix string, flag int) *log.Logger = log.New

// NewFileLogger
/*
@param logPath 	(1) 如果文件不存在，会生成文件和父目录
				(2) 文件存在，新的内容会append
*/
func NewFileLogger(filePath, prefix string, perm os.FileMode) (*log.Logger, error) {
	if err := fileKit.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}

	/* out */
	// 如果生成文件的话，其权限为0666（与os.Create()的权限一样）
	// os.OpenFile()的传参flag，可参考：https://studygolang.com/articles/22180
	f, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, perm)
	if err != nil {
		return nil, err
	}
	// 此处不能关闭writer，否则日志内容将写不进去
	//defer f.Close()

	flag := log.Ldate | log.Ltime | log.Lmicroseconds | log.Lmsgprefix
	return NewLogger(f, prefix, flag), nil
}
