package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

const (
	// AllPerm 所有权限（0777）
	AllPerm = os.ModePerm

	// SafePerm 文件拥有者可以读写文件，而组用户和其他用户只能读取文件.
	/*
		3种权限:
		(1) 4: 读权限
		(2) 2: 写权限
		(3) 1: 执行权限

		第1位: 表示特殊权限位（不常用，通常为 0）
		第2位: 表示文件所有者的权限
		第3位: 表示所属组的权限
		第4位: 表示其他用户的权限
	*/
	SafePerm os.FileMode = 0644
)

// IsReadable 是否有 读 权限?
/*
	@param path 文件（或目录）的路径
	@return 传参path不存在的话，将返回false

	e.g. ("") => false
*/
func IsReadable(path string) bool {
	return gfile.IsReadable(path)
}

// IsWritable 是否有 写 权限?
/*
	@param path 文件（或目录）的路径
	@return 传参path不存在的话，将返回false

	e.g. ("") => false
*/
func IsWritable(path string) bool {
	return gfile.IsWritable(path)
}

// GetFileMode get mode and permission bits of file/directory
func GetFileMode(path string) (os.FileMode, error) {
	if err := AssertExist(path); err != nil {
		return 0, err
	}

	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Mode(), nil
}
