package main

import (
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"fmt"
)

func main() {
	fmt.Println(idKit.NewXid())
}
