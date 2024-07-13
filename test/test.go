package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/randomKit"
)

func main() {
	for i := range 100000 {
		i = i

		tmp := randomKit.Int(0, 10)
		if tmp < 0 || tmp >= 10 {
			panic(tmp)
		}
		fmt.Println(tmp)
	}
}
