package fileKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"github.com/richelieu-yang/chimera/v2/src/funcKit"
)

// AssertExist
/*
@param path 文件（或目录）的路径
*/
func AssertExist(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertNotExistOrIsFile
/*
通过的情况: 	不存在 || 存在但是个文件
不通过的情况:	存在但是个目录

@param mkdirArgs true: path不存在的话，为其创建父目录（可多级）
*/
func AssertNotExistOrIsFile(path string, mkdirArgs ...bool) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}
	if Exists(path) && IsDir(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}

	var flag bool = true
	if mkdirArgs != nil {
		flag = mkdirArgs[0]
	}
	if flag {
		if err := MkParentDirs(path); err != nil {
			return err
		}
	}

	return nil
}

// AssertNotExistOrIsDir
/*
通过的情况: 	不存在 || 存在但是个目录
不通过的情况:	存在但是个文件

@param mkdirArgs true: path不存在的话，为其创建目录（可多级）
*/
func AssertNotExistOrIsDir(path string, mkdirArgs ...bool) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}
	if Exists(path) && IsFile(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}

	var flag bool = true
	if mkdirArgs != nil {
		flag = mkdirArgs[0]
	}
	if flag {
		if err := MkDirs(path); err != nil {
			return err
		}
	}

	return nil
}

// AssertExistAndIsFile
/*
@return 如果path存在且是个文件，返回nil
*/
func AssertExistAndIsFile(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsDir(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) exists but it is a directory", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertExistAndIsDir
/*
@return 如果path存在且是个目录，返回nil
*/
func AssertExistAndIsDir(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if IsFile(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) exists but it is a file", funcKit.GetFuncName(1), path)
	}
	return nil
}

// AssertReadableAndWritable
/*
前提: path 非blank && 存在.
*/
func AssertReadableAndWritable(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsReadable(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) isn't readable", funcKit.GetFuncName(1), path)
	}
	if !IsWritable(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) isn't writable", funcKit.GetFuncName(1), path)
	}
	return nil
}

func AssertReadable(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsReadable(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) isn't readable", funcKit.GetFuncName(1), path)
	}
	return nil
}

func AssertWritable(path string) error {
	if strKit.IsBlank(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) is blank", funcKit.GetFuncName(1), path)
	}

	if !Exists(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) doesn't exist", funcKit.GetFuncName(1), path)
	}
	if !IsWritable(path) {
		return errorKit.NewSkip(1, "[%s] path(%s) isn't writable", funcKit.GetFuncName(1), path)
	}
	return nil
}
