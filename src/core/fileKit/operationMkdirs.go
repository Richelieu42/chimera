package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
	"os"
)

func MkDirs(dirPaths ...string) error {
	return MkDirsWithPerm(os.ModePerm, dirPaths...)
}

func MkParentDirs(paths ...string) error {
	return MkParentDirsWithPerm(os.ModePerm, paths...)
}

// MkDirsWithPerm 为目录路径，创建（一级或多级）目录.
/*
PS:
(1) 如果目录已经存在，将返回nil；
(2) 如果 传参dirPath 对应的是个已存在的文件，将返回error（"mkdir {xxx}: not a directory"）.

@param dirPaths 目录路径s（相对路径 || 绝对路径）

e.g.
("i:/test/test.exe") 	=> 	路径没问题且目录不存在的情况下，会在i盘创建"test"、"test.exe"两个目录
("i:/test1/test2/")		=>	路径没问题且目录不存在的情况下，会在i盘创建"test1"、"test2"两个目录

e.g.1 Mac
("")					=>	nil（什么都不会做）
("/")					=>	nil（什么都不会做）
(".")					=>	nil（什么都不会做）
("./")					=>	nil（什么都不会做）
*/
func MkDirsWithPerm(perm os.FileMode, dirPaths ...string) error {
	for _, dirPath := range dirPaths {
		// os.MkdirAll() 的第一个传参
		// (1) 如果为""会返回error(mkdir : no such file or directory)
		// (2) 如果为多个空格，返回的error为nil（并不会创建目录）
		if strKit.IsEmpty(dirPath) {
			continue
		}

		if err := os.MkdirAll(dirPath, perm); err != nil {
			err = errorKit.Wrap(err, `fail with dirPath(%s) and perm(%d)`, dirPath, perm)
			return err
		}
	}
	return nil
}

// MkParentDirsWithPerm 为父路径，创建（一级或多级）目录.
/*
@param filePaths （文件 || 目录）路径s（相对路径 || 绝对路径）

e.g.
("")	=> nil
(".")	=> nil
*/
func MkParentDirsWithPerm(perm os.FileMode, paths ...string) error {
	for _, path := range paths {
		// Richelieu: 为防止 import cycle，不直接使用 pathKit.ParentDir
		parentDir := gfile.Dir(path)
		if err := MkDirsWithPerm(perm, parentDir); err != nil {
			return err
		}
	}
	return nil
}
