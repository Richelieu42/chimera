package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
)

func main() {
	fmt.Println(excelKit.CoordinatesToCellName(1, 1))       // A1 <nil>
	fmt.Println(excelKit.CoordinatesToCellName(1, 1, true)) // $A$1 <nil>
}
