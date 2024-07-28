package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/atomic/atomicKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
)

func main() {
	netKit.Dial

	i := atomicKit.NewInt64(0)
	fmt.Println(i.Inc() - 1)
	fmt.Println(i.Inc() - 1)
	fmt.Println(i.Inc() - 1)
}
