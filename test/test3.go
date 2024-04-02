package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"io"
	"time"
)

func main() {
	path := "/Users/richelieu/Documents/ino.7z"

	f, err := fileKit.Open(path)
	if err != nil {
		panic(err)
	}

	start := time.Now()

	var reader io.Reader
	reader = f
	//reader = bufio.NewReader(f)
	_, err = ioKit.ReadAll(reader)
	if err != nil {
		panic(err)
	}

	fmt.Println(time.Since(start))
}
