package embedKit

import (
	"embed"
	"io/fs"
	"net/http"
)

func NewHttpFileSystem1(embedFs embed.FS, dir string) (http.FileSystem, error) {
	//subFs, err := fs.Sub(embedFs, dir)
	//if err != nil {
	//	return nil, err
	//}
	return http.FS(embedFs), nil
}

// NewHttpFileSystem
/*
@param dir 子目录
*/
func NewHttpFileSystem(embedFs embed.FS, dir string) (http.FileSystem, error) {
	subFs, err := fs.Sub(embedFs, dir)
	if err != nil {
		return nil, err
	}
	return http.FS(subFs), nil
}
