package excelKit

import "github.com/xuri/excelize/v2"

// SetCellValue 设置单元格的值.
/*
@param sheetName 	工作表名
@param col 			[1, 16384]
@param row 			[1, 1048576]
*/
func SetCellValue(f *excelize.File, sheetName string, col, row int, value interface{}) error {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return err
	}
	return f.SetCellValue(sheetName, cellName, value)
}
