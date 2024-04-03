package excelKit

import "github.com/xuri/excelize/v2"

func GetCellType(f *excelize.File, sheetName string, col, row int) (excelize.CellType, error) {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return 0, err
	}
	return f.GetCellType(sheetName, cellName)
}
