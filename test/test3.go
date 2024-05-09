package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
)

func main() {
	fmt.Println(dataSizeKit.ToReadableIecString(1048576))
}
