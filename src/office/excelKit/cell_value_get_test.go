package excelKit

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestGetCellValue(t *testing.T) {
	path := "_test.xlsx"

	f, err := OpenFile(path)
	if err != nil {
		panic(err)
	}
	defer func() {
		//if err := f.Save(); err != nil {
		//	logrus.WithError(err).Error("Fail to save.")
		//	return
		//}
		if err := f.Close(); err != nil {
			logrus.WithError(err).Error("Fail to close.")
		}
	}()

	activeSheetName := GetActiveSheetName(f)
	logrus.Infof("activeSheetName: %s", activeSheetName)
	fmt.Println(GetCellValue(f, activeSheetName, 3, 3))
	fmt.Println(GetCellValue(f, activeSheetName, 3, 4))
	fmt.Println(GetCellValue(f, activeSheetName, 3, 5))
}
