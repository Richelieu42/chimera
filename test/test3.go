package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
)

func main() {
	fmt.Println(excelKit.JoinCellName("C", 3)) // C3 <nil>

	fmt.Println(excelKit.SplitCellName("C3")) // C 3 <nil>
	fmt.Println(excelKit.SplitCellName("c3")) // c 3 <nil>
}
