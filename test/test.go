package main

import (
	"fmt"
	"github.com/klauspost/cpuid/v2"
	"github.com/richelieu-yang/chimera/v3/src/core/cpuKit"
)

func main() {
	fmt.Println(cpuKit.HasFeature(cpuid.AVX))
}
