package main

import (
	"github.com/goccy/go-json"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"

	"fmt"
)

func main() {
	json.Marshal()

	fmt.Println(idKit.NewXid())
}
