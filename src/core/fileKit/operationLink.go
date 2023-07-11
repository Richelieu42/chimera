package fileKit

import "os"

var (
	// CreateSoftLink 创建软链接
	/*
	   Golang 之 文件硬连接 与 软连接: https://blog.csdn.net/icebergliu1234/article/details/109208030

	   @param src	源文件
	   @param dest	生成链接的位置
	*/
	CreateSoftLink func(oldname, newname string) error = os.Symlink

	// CreateHardLink 创建软链接
	/*
	   Golang 之 文件硬连接 与 软连接: https://blog.csdn.net/icebergliu1234/article/details/109208030

	   @param src	源文件
	   @param dest	生成链接的位置
	*/
	CreateHardLink func(oldname, newname string) error = os.Link
)
