package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
)

func main() {
	fmt.Println(dataSizeKit.ToReadableIecString(314572800))

	//fmt.Println(dataSizeKit.MiB * 300)

	//ctx := context.WithValue(context.TODO(), "a", "A")
	//ctx = context.WithValue(ctx, "b", true)
	//ctx = context.WithValue(ctx, "c", 996)
	//
	//fmt.Println(ctx.Value("a")) // A
	//fmt.Println(ctx.Value("b")) // true
	//fmt.Println(ctx.Value("c")) // 996
}
