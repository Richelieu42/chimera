package embedKit

import (
	"embed"
	"net/http"
	"testing"
)

//go:embed c/*
var tmpEfs embed.FS

func TestNewHttpFileSystem(t *testing.T) {
	httpFs, err := NewHttpFileSystem(tmpEfs, "c")
	if err != nil {
		panic(err)
	}

	// TODO: 将 "/" 改成 "/s"，会有问题
	http.Handle("/", http.FileServer(httpFs))
	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
