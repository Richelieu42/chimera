package main

import (
	"fmt"
	"github.com/richelieu42/go-scales/src/http/httpKit"
)

type MyInt int

type Bean struct {
}

func main() {
	header := make(map[string][]string)

	// add
	httpKit.AddHeader(header, "k", "0")
	httpKit.AddHeader(header, "k", "1")
	fmt.Println(header) // map[K:[0 1]]

	// get
	fmt.Println(httpKit.GetHeaderValues(header, "k"))  // 0
	fmt.Println(httpKit.GetHeaderValues(header, "k1")) // 0
}
