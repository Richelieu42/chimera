package main

import (
	"embed"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"io/fs"
	"net/http"
)

//go:embed c
var efs embed.FS

func main() {
	handler, err := NewHttpHandler(efs, "c")
	if err != nil {
		panic(err)
	}

	// TODO: 将 "/" 改成 "/s"，会有问题
	http.Handle("/", handler)
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}

func NewHttpHandler(embedFs embed.FS, dir string) (http.Handler, error) {
	if err := strKit.AssertNotEmpty(dir, "dir"); err != nil {
		return nil, err
	}

	subFs, err := fs.Sub(embedFs, dir)
	if err != nil {
		return nil, err
	}
	httpFs := http.FS(subFs)
	return http.FileServer(httpFs), nil
}
