package excelKit

import "github.com/xuri/excelize/v2"

func GetActiveSheetName(f *excelize.File) string {
	i := f.GetActiveSheetIndex()
	return f.GetSheetName(i)
}
