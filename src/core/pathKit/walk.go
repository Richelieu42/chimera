package pathKit

import (
	"io/fs"
	"path/filepath"
)

// Walk 遍历目录.
/*
Deprecated: 使用 WalkDir.

PS:
(1) 遍历包含 传参root;
(2) filepath.Walk 和 filepath.WalkDir 的区别:
	(a) filepath.WalkDir 是在 Go1.16 中引入的，比 filepath.Walk 更高效，filepath.WalkDir 避免了对每个访问的文件或目录调用 os.Lstat
	(b) 回调函数fn的传参不同，分别为 filepath.WalkFunc 和 fs.WalkDirFunc，
		此外，WalkDirFunc 在读取目录之前调用，以允许 SkipDir 完全跳过目录读取。如果目录读取失败，则该函数会再次为该目录调用一次以报告错误.
*/
var Walk func(root string, fn filepath.WalkFunc) error = filepath.Walk

var WalkDir func(root string, fn fs.WalkDirFunc) error = filepath.WalkDir
