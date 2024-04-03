package main

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
)

func main() {
	path := "/Users/richelieu/Desktop/未命名.xlsx"
	f, err := excelKit.NewFileWithPath(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Close())

	fileutil.FileSize()

}
