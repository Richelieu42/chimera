package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
)

func main() {
	path := "/Users/richelieu/Desktop/未命名.xlsx"

	{
		size, err := fileKit.GetSize(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(dataSizeKit.ToReadableSiString(float64(size)))
	}

	f, err := excelKit.NewFileWithPath(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(f.Close())

	{
		size, err := fileKit.GetSize(path)
		if err != nil {
			panic(err)
		}
		fmt.Println(dataSizeKit.ToReadableSiString(float64(size)))
	}
}
