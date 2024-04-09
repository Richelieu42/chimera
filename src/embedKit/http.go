package embedKit

import (
	"embed"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"io/fs"
	"net/http"
)

// NewHttpFileSystem
/*
@param dir 子目录
*/
func NewHttpFileSystem(embedFs embed.FS, dir string) (http.FileSystem, error) {
	if err := strKit.AssertNotEmpty(dir, "dir"); err != nil {
		return nil, err
	}

	subFs, err := fs.Sub(embedFs, dir)
	if err != nil {
		return nil, err
	}
	return http.FS(subFs), nil
}
