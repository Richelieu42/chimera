package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/ip/ipRegionKit"
)

func main() {
	fmt.Println(ipRegionKit.GetRegion("180.98.201.169"))
}
