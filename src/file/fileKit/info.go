package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"os"
	"path/filepath"
)

var (
	Exists func(path string) bool = gfile.Exists
	IsFile func(path string) bool = gfile.IsFile
	IsDir  func(path string) bool = gfile.IsDir

	// Stat 获取文件（或目录）信息
	/*
		@param path 如果为""或不存在，将返回error(e.g."" => stat : no such file or directory)
	*/
	Stat func(path string) (os.FileInfo, error) = gfile.Stat

	// IsEmpty checks whether the given `path` is empty.
	/*
		If `path` is a folder, it checks if there's any file under it.
		If `path` is a file, it checks if the file size is zero.
		Note that it returns true if `path` does not exist.
	*/
	IsEmpty func(path string) bool = gfile.IsEmpty

	// GetFileName 获取 文件名.
	/*
		e.g.
		/var/www/file.js -> file.js
		file.js          -> file.js
	*/
	GetFileName func(path string) string = filepath.Base

	// GetName 获取 文件名的前缀.
	/*
		e.g.
		/var/www/file.js -> file
		file.js          -> file
	*/
	GetName func(path string) string = gfile.Name

	// GetExt 获取 文件名的后缀（带"."）
	/*
		@return 可能为""

		e.g.
			println(fileKit.GetExt("main.go"))  // ".go"
			println(fileKit.GetExt("api.json")) // ".json"
			println(fileKit.GetExt(""))         // ""
			println(fileKit.GetExt("    "))     // ""
			println(fileKit.GetExt("empty"))    // ""
	*/
	GetExt func(path string) string = gfile.Ext

	// GetExtName 获取 后缀（不带"."）
	/*
		@return 可能为""
		e.g.
			println(fileKit.GetExtName("main.go"))  // "go"
			println(fileKit.GetExtName("api.json")) // "json"
			println(fileKit.GetExtName(""))         // ""
			println(fileKit.GetExtName("    "))     // ""
			println(fileKit.GetExtName("empty"))    // ""
	*/
	GetExtName func(path string) string = gfile.ExtName
)
