package embedKit

import (
	"embed"
	"io/fs"
	"net/http"
)

func ToHttpFileSystem(embedFs embed.FS) http.FileSystem {
	return http.FS(embedFs)
}

// Sub 获取一个embed.FS实例的子embed.FS实例.
func Sub(fsys fs.FS, dir string) (fs.FS, error) {
	return fs.Sub(fsys, dir)
}

// NewHttpFileSystem
/*
@param dir 子目录
*/
func NewHttpFileSystem(embedFs embed.FS, dir string) (httpFs http.FileSystem, err error) {
	subFs, err := fs.Sub(embedFs, dir)
	if err != nil {
		return nil, err
	}
	return http.FS(subFs), nil
}
