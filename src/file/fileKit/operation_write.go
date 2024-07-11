package fileKit

import (
	"github.com/duke-git/lancet/v2/fileutil"
	"os"
)

// WriteToFile 将数据（[]byte, 字节流）写到文件中.
/*
@param filePath 目标文件的路径
				(1) 不存在的话，会创建一个新的文件;
				(2) 存在且是个文件的话，会 "覆盖" 掉旧的（并不会加到该文件的最后面）.
*/
func WriteToFile(filePath string, data []byte, perm os.FileMode) error {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := MkParentDirs(filePath); err != nil {
		return err
	}

	return os.WriteFile(filePath, data, perm)
}

// WriteStringToFile 将数据（字符串）写到文件中.
/*
@param filePath 目标文件的路径，	(1)不存在的话，会创建一个新的文件;
							 	(2) 存在且是个文件的话，由 传参append 决定.
*/
func WriteStringToFile(filePath string, content string, append bool) error {
	if err := AssertNotExistOrIsFile(filePath); err != nil {
		return err
	}
	if err := MkParentDirs(filePath); err != nil {
		return err
	}

	return fileutil.WriteStringToFile(filePath, content, append)
}
