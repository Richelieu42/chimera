package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
)

func main() {
	fmt.Print(netKit.JoinToHost("::1", 8888))
}
