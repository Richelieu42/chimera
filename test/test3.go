package main

import (
	"fmt"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
)

func main() {
	fmt.Println(excelKit.ColumnNameToNumber("A"))  // 1 <nil>
	fmt.Println(excelKit.ColumnNameToNumber("a"))  // 1 <nil>
	fmt.Println(excelKit.ColumnNameToNumber("B"))  // 2 <nil>
	fmt.Println(excelKit.ColumnNameToNumber("AK")) // 37 <nil>
}
