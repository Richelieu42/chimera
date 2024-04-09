package main

import (
	"embed"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"io/fs"
	"net/http"
)

//go:embed c/*
var efs embed.FS

func main() {
	httpFs, err := NewHttpFileSystem(efs, "c")
	if err != nil {
		panic(err)
	}

	// TODO: 将 "/" 改成 "/s"，会有问题
	http.Handle("/", http.FileServer(httpFs))
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

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
