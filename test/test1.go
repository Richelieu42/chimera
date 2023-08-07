package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/core/memoryKit"
	"github.com/richelieu-yang/chimera/v2/src/dataSizeKit"
)

func main() {
	stats := memoryKit.GetProgramMemoryStats()

	fmt.Println("Alloc:", dataSizeKit.ToReadableStringWithIEC(stats.Alloc))
	fmt.Println("TotalAlloc:", dataSizeKit.ToReadableStringWithIEC(stats.TotalAlloc))
	fmt.Println("Sys:", dataSizeKit.ToReadableStringWithIEC(stats.Sys))
	fmt.Println("NumGC:", stats.NumGC)

	fmt.Println(stats.EnableGC)
	fmt.Println(stats.DebugGC)

	fmt.Println(stats)
}
