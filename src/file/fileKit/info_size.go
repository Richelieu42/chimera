package fileKit

import (
	"github.com/duke-git/lancet/v2/fileutil"
)

var (
	GetFileSize func(filePath string) (int64, error) = fileutil.FileSize

	GetDirSize func(dirPath string) (int64, error) = fileutil.DirSize
)

// GetSize 获取文件（或目录）的大小（字节数）.
/*
@param path 文件路径 || 目录路径
*/
func GetSize(path string) (int64, error) {
	if err := AssertExist(path); err != nil {
		return 0, err
	}

	if IsDir(path) {
		return GetDirSize(path)
	}
	return GetFileSize(path)
}

//// GetSize 获取文件（或目录）的大小.
//func GetSize(path string) (int64, error) {
//	if err := AssertExist(path); err != nil {
//		return 0, err
//	}
//
//	if IsFile(path) {
//		return getFileSize(path)
//	}
//	return getDirSize(path)
//}
//
//// getFileSize 获取文件的大小.
//func getFileSize(filePath string) (int64, error) {
//	info, err := os.Stat(filePath)
//	if err != nil {
//		return 0, err
//	}
//	return info.Size(), nil
//}
//
//// getDirSize 获取目录的大小（包含其内文件和目录）.
///*
//参考:
//golang获取文件/目录（包含下面的文件）的大小
//	https://blog.csdn.net/n_fly/article/details/117080173
//*/
//func getDirSize(dirPath string) (int64, error) {
//	var bytes int64
//	err := filepath.Walk(dirPath, func(_ string, info os.FileInfo, err error) error {
//		if !info.IsDir() {
//			bytes += info.Size()
//		}
//		// 如果 err != nil，将中止遍历
//		return err
//	})
//	if err != nil {
//		return 0, err
//	}
//	return bytes, nil
//}
