package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
)

func main() {
	fmt.Println(netKit.AssertHost("127.0.0.1:100001"))
}
