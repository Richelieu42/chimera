package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"

	"fmt"
)

func main() {
	fmt.Println(jsonKit.GetLibrary())

	//json.Marshal()
	//fmt.Println(idKit.NewXid())
}
