package main

import (
	"github.com/playwright-community/playwright-go"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"

	"fmt"
)

func main() {
	fmt.Println(jsonKit.GetLibrary())

	playwright.String()

	//json.Marshal()
	//fmt.Println(idKit.NewXid())
}
