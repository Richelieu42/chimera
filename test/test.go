package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
)

func main() {
	path := "./iShot_2024-09-04_13.51.58.PNG"

	fmt.Println(fileKit.GetExt(path))
	fmt.Println(fileKit.GetExtName(path))
}
