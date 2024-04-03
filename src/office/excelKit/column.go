package excelKit

import "github.com/xuri/excelize/v2"

var (
	// ColumnNumberToName 列数字 => 列名
	/*
		@param num 有效范围: [1, 16384]

		e.g.
			fmt.Println(excelKit.ColumnNumberToName(1))  // A <nil>
			fmt.Println(excelKit.ColumnNumberToName(2))  // B <nil>
			fmt.Println(excelKit.ColumnNumberToName(10)) // J <nil>
			fmt.Println(excelKit.ColumnNumberToName(37)) // AK <nil>
	*/
	ColumnNumberToName func(num int) (string, error) = excelize.ColumnNumberToName

	// ColumnNameToNumber 列名 => 列数字
	/*
		e.g.
			fmt.Println(excelKit.ColumnNameToNumber("A"))  // 1 <nil>
			fmt.Println(excelKit.ColumnNameToNumber("a"))  // 1 <nil>
			fmt.Println(excelKit.ColumnNameToNumber("B"))  // 2 <nil>
			fmt.Println(excelKit.ColumnNameToNumber("AK")) // 37 <nil>
	*/
	ColumnNameToNumber func(name string) (int, error) = excelize.ColumnNameToNumber
)
