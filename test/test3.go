package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"os"
)

func main() {
	path := "/Users/richelieu/Documents/ino.7z"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	_, err = io.ReadAll(f)
	if err != nil {
		panic(err)
	}
}
