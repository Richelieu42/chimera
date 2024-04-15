package embedKit

import (
	"embed"
	"io/fs"
	"net/http"
)

func ToHttpFileSystem(embedFs embed.FS) http.FileSystem {
	return http.FS(embedFs)
}

// NewHttpFileSystem
/*
@param dir 子目录
*/
func NewHttpFileSystem(embedFs embed.FS, dir string) (httpFs http.FileSystem, err error) {
	var subFs fs.FS
	subFs, err = fs.Sub(embedFs, dir)
	if err != nil {
		return
	}

	httpFs = http.FS(subFs)
	return
}
