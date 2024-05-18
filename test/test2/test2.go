package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
)

func main() {
	fmt.Println(urlKit.PolyfillUrl("http://example.com/users/{userId}", nil))
}
