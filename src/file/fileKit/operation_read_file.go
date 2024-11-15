package fileKit

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

// ReadFile 读取文件的数据.
/*
PS:
(1) ioutil.ReadFile() 比 ioutil.ReadAll() 性能好，特别是大文件；
(2) 编码必须为"UTF-8"！！！
(3) 读取大文件也很快.

@param path 文件的路径（不能是目录的路径）
*/
func ReadFile(filePath string) ([]byte, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	return os.ReadFile(filePath)
}

func ReadFileToString(filePath string) (string, error) {
	data, err := ReadFile(filePath)
	return string(data), err
}

func ReadFileByLine(filePath string) ([]string, error) {
	if err := AssertExistAndIsFile(filePath); err != nil {
		return nil, err
	}

	return fileutil.ReadFileByLine(filePath)
}

func ReadLines(file string, callback func(line string) error) error {
	return gfile.ReadLines(file, callback)
}

func ReadLinesBytes(file string, callback func(bytes []byte) error) error {
	return gfile.ReadLinesBytes(file, callback)
}

//var (
//	ReadLines func(file string, callback func(line string) error) error = gfile.ReadLines
//
//	ReadLinesBytes func(file string, callback func(bytes []byte) error) error = gfile.ReadLinesBytes
//)
