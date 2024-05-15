package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
)

func main() {
	fmt.Println(urlKit.EncodeURIComponent(":$"))
}
