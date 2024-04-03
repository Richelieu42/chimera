package main

import (
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/office/excelKit"
	"github.com/sirupsen/logrus"
)

func main() {
	path := "/Users/richelieu/Desktop/未命名.xlsx"

	f, err := excelKit.NewFileWithPath(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := f.Save(); err != nil {
			logrus.WithError(err).Error("Fail to save.")
			return
		}
		if err := f.Close(); err != nil {
			logrus.WithError(err).Error("Fail to close.")
		}
	}()

	sheetName := f.GetSheetName(f.GetActiveSheetIndex())
	if err := excelKit.SetCellStr(f, sheetName, 3, 3, "hello"); err != nil {
		panic(err)
	}
	if err := excelKit.SetCellStr(f, sheetName, 3, 4, "world"); err != nil {
		panic(err)
	}
	if err := excelKit.SetCellValue(f, sheetName, 3, 5, "!"); err != nil {
		panic(err)
	}
}
