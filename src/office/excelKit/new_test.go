package excelKit

import (
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestNewFile(t *testing.T) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")
	if err != nil {
		panic(err)
	}

	// Set value of a cell.
	if err := f.SetCellValue("Sheet2", "A2", "Hello world."); err != nil {
		panic(err)
	}
	if err := f.SetCellValue("Sheet1", "B2", 100); err != nil {
		panic(err)
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet by the given path.
	if err := f.SaveAs("_test.xlsx"); err != nil {
		panic(err)
	}
}
