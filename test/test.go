package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
)

func main() {
	s := []byte("{}[]")
	for _, v := range s {
		fmt.Println(string(v), v)
	}
}

func IsJsonString(str string) (rst bool) {
	if strKit.StartWith(str, "{") {
		rst = strKit.EndWith(str, "}")
	} else if strKit.StartWith(str, "[") {
		rst = strKit.EndWith(str, "]")
	}
	return
}
