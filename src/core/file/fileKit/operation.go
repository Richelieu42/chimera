package fileKit

import (
	"github.com/richelieu42/chimera/v2/src/assert/fileAssert"
	"os"
	"path/filepath"
)

// NewTemporaryFile 在指定目录下，生成临时文件.
/*
@param dirPath 如果为""，临时文件将生成在 系统临时目录 内；如果为"."，临时文件将生成在 当前目录 内.

e.g.
pattern: "tempfile_test" 		=> 临时文件的文件名: "tempfile_test2594316144"
pattern: "tempfile_test*" 		=> 临时文件的文件名: "tempfile_test827818253"
pattern: "tempfile_test*.xyz" 	=> 临时文件的文件名: "tempfile_test3617672388.xyz"
*/
func NewTemporaryFile(dirPath, pattern string) (*os.File, error) {
	// check dirPath
	if err := fileAssert.AssertNotExistOrIsDir(dirPath); err != nil {
		return nil, err
	}
	if err := MkDirs(dirPath); err != nil {
		return nil, err
	}

	return os.CreateTemp(dirPath, pattern)
}

// NewFile 创建文件.
/*
PS: 如果文件已经存在，会覆盖掉它.
*/
func NewFile(filePath string) (*os.File, error) {
	// 检查 filePath
	if err := fileAssert.AssertNotExistOrIsFile(filePath); err != nil {
		return nil, err
	}
	if err := MkParentDirs(filePath); err != nil {
		return nil, err
	}

	return os.Create(filePath)
}

// WriteToFile 将数据（字节流）写到文件中.
/*
@param target 目标文件的路径（不存在的话，会创建一个新的文件；存在且是个文件的话，会覆盖掉旧的（并不会加到该文件的最后面））
*/
func WriteToFile(data []byte, dest string) error {
	// 检查 dest
	if err := fileAssert.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := MkParentDirs(dest); err != nil {
		return err
	}

	return os.WriteFile(dest, data, os.ModePerm)
}

// Delete 删除 文件 或 目录（内部有文件或目录，也会一并删除）.
/*
@param path 文件（或目录）的路径（绝对 || 相对），可以不存在，此时将返回nil

PS:
(1) 传参path可以为"": 正常执行，返回nil；
(2) path对应的文件或目录不存在: 正常执行，返回nil；
(3) Windows系统，如果 传参path 对应的是 一个被锁定的文件 或者 一个目录（内部有文件被锁定），将返回error（remove xxx(path): The process cannot access the file because it is being used by another process.）.
*/
func Delete(path string) error {
	return os.RemoveAll(path)
}

// EmptyDir 清空目录：删掉目录中的文件和子目录（递归），但该目录本身不会被删掉.
/*
@param dirPath 可以不存在，此时将返回nil
*/
func EmptyDir(dirPath string) error {
	if NotExist(dirPath) {
		return nil
	}
	if err := fileAssert.AssertExistAndIsDir(dirPath); err != nil {
		return err
	}

	// 遍历目录
	dirEntries, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}
	for _, dirEntry := range dirEntries {
		if err := os.RemoveAll(filepath.Join(dirPath, dirEntry.Name())); err != nil {
			return err
		}
	}
	return nil
}
