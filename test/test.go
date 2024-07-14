package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
)

func main() {
	for i := range 100 {
		i = i
		fmt.Println(randomKit.RandFloat(1, 2, 3))
	}
}
