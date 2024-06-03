package embedKit

import (
	"embed"
	"io/fs"
	"net/http"
)

var (
	// Sub 获取一个embed.FS实例的子embed.FS实例.
	Sub func(fsys fs.FS, dir string) (fs.FS, error) = fs.Sub
)

// NewSubHttpFileSystem
/*
@param dir 子目录
*/
func NewSubHttpFileSystem(embedFs embed.FS, dir string) (httpFs http.FileSystem, err error) {
	subFs, err := fs.Sub(embedFs, dir)
	if err != nil {
		return nil, err
	}
	return http.FS(subFs), nil
}
