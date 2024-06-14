package main

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/dataSizeKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
)

func main() {

	fmt.Println(timeKit.FormatCurrent("2006-01-02T15:04:05.000Z07:00 MST"))

	fmt.Println(timeKit.FormatCurrent("2006-01-02 15:04:05 -0700 MST"))

	fmt.Println(dataSizeKit.MiB * 512)
}
