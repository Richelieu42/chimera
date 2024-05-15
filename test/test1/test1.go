package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
)

func main() {
	u, err := urlKit.Parse("http://127.0.0.1:80")
	if err != nil {
		panic(err)
	}
	fmt.Println(u)
}
