package excelKit

import (
	"github.com/xuri/excelize/v2"
)

func GetCellValue(f *excelize.File, sheetName string, col, row int, opts ...excelize.Options) (string, error) {
	cellName, err := CoordinatesToCellName(col, row)
	if err != nil {
		return "", err
	}
	return f.GetCellValue(sheetName, cellName, opts...)
}
