package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"time"
)

func main() {
	path := "nohup.out"

	start := time.Now()
	if err := fileKit.CopyFile(path, "nohup1.out"); err != nil {
		panic(err)
	}
	fmt.Println(time.Since(start))
	if _, err := fileKit.Create(path); err != nil {
		panic(err)
	}
}
