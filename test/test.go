package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/mathKit"
)

func main() {
	fmt.Println(mathKit.Round(3.1029, 1)) // 3.1
	fmt.Println(mathKit.Round(3.1029, 2)) // 3.1
	fmt.Println(mathKit.Round(3.1029, 3)) // 3.103
}
