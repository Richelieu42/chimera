package excelKit

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetCellValue(t *testing.T) {
	path := "_test.xlsx"

	f, err := NewFileWithPath(path)
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

	activeSheetName := GetActiveSheetName(f)
	logrus.Infof("activeSheetName: %s", activeSheetName)
	if err := SetCellStr(f, activeSheetName, 3, 3, "hello"); err != nil {
		panic(err)
	}
	if err := SetCellStr(f, activeSheetName, 3, 4, "world"); err != nil {
		panic(err)
	}
	if err := SetCellValue(f, activeSheetName, 3, 5, "!"); err != nil {
		panic(err)
	}
}
