package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/image/imageKit"
	"os"
)

func main() {
	path := "/Users/richelieu/Desktop/iShot_2024-09-04_13.51.58.PNG"

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, tmp, err := imageKit.Decode(f)
	if err != nil {
		panic(err)
	}
	fmt.Println(tmp)
	fmt.Println(img.Bounds().String())
}
