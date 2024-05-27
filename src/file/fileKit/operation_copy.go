package fileKit

import (
	"github.com/gogf/gf/v2/os/gfile"
)

var (
	// Copy 复制文件（或目录）
	/*
		PS: 可以设置权限.
	*/
	Copy func(src string, dst string, option ...gfile.CopyOption) error = gfile.Copy

	// CopyFile 复制文件
	CopyFile func(src, dst string, option ...gfile.CopyOption) (err error) = gfile.CopyFile

	// CopyDir 复制目录
	CopyDir func(src string, dst string, option ...gfile.CopyOption) (err error) = gfile.CopyDir
)

//// CopyFile 复制单个文件.
///*
//@param src	一个已经存在的文件
//@param dest	一个已经存在的文件（会覆盖） || 一个不存在的文件
//@return 第一个返回值: the number of bytes copied（单位为byte）
//*/
//func CopyFile(src, dest string) (int64, error) {
//	// 检查 src
//	if err := AssertExistAndIsFile(src); err != nil {
//		return 0, err
//	}
//	// 检查 dest
//	if err := AssertNotExistOrIsFile(dest); err != nil {
//		return 0, err
//	}
//
//	if err := MkParentDirs(dest); err != nil {
//		return 0, err
//	}
//	srcFile, err := os.OpenReadOnly(src)
//	if err != nil {
//		return 0, err
//	}
//	defer srcFile.Close()
//
//	destFile, err := os.Create(dest)
//	if err != nil {
//		return 0, err
//	}
//	defer destFile.Close()
//
//	return io.Copy(destFile, srcFile)
//}
//
//// CopyDir 将 src目录的所有内容 复制到 dest目录 中.
///*
//PS:
//(1) src目录下如果还有目录，会递归（空目录也会复制过去）；
//(2) 类似于Linux的 cp -r 命令.
//
//@param src	一个已经存在的目录
//@param dest	一个已经存在的目录 || 一个不存在的目录
//*/
//func CopyDir(src, dest string) error {
//	// 检查 src
//	if err := AssertExistAndIsDir(src); err != nil {
//		return err
//	}
//	// 检查 dest
//	if err := AssertNotExistOrIsDir(dest); err != nil {
//		return err
//	}
//
//	if err := MkDirs(dest); err != nil {
//		return err
//	}
//	// 遍历
//	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//
//		relPath, err := filepath.Rel(src, path)
//		if err != nil {
//			return err
//		}
//		// 此次跳过
//		if relPath == "." {
//			return nil
//		}
//		tmpSrc := filepath.Join(src, relPath)
//		tmpDest := filepath.Join(dest, relPath)
//		if info.IsDir() {
//			return CopyDir(tmpSrc, tmpDest)
//		}
//		_, err = CopyFile(tmpSrc, tmpDest)
//		return err
//	})
//}
