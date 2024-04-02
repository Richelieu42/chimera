package main

import (
	"bytes"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
)

func main() {
	buf := bytes.NewBuffer(nil)
	reader := ioKit.NewReaderFromString("abcdefg")

	fmt.Println(ioKit.Copy(buf, reader)) // 7 <nil>
	data, err := ioKit.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	fmt.Println(len(data)) // 0（因为Copy操作会"读"reader）
}
