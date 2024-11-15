package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
)

const (
	// AllPerm 所有权限
	AllPerm = os.ModePerm

	// SafePerm 文件拥有者可以读写文件，而组用户和其他用户只能读取文件.
	/*
		4: 读权限
		2: 写权限
		1: 执行权限
	*/
	SafePerm os.FileMode = 0644
)

var (
	// IsReadable 是否有 读 权限?
	/*
		@param path 文件（或目录）的路径
		@return 传参path不存在的话，将返回false

		e.g.
			("") => false
	*/
	IsReadable func(path string) bool = gfile.IsReadable

	// IsWritable 是否有 写 权限?
	/*
		@param path 文件（或目录）的路径
		@return 传参path不存在的话，将返回false

		e.g.
			("") => false
	*/
	IsWritable func(path string) bool = gfile.IsWritable
)

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
