package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
)

func main() {
	i := atomicKit.NewInt64(0)
	fmt.Println(i.Inc() - 1)
	fmt.Println(i.Inc() - 1)
	fmt.Println(i.Inc() - 1)
}
