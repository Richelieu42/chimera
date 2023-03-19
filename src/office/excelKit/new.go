package excelKit

import (
	"github.com/richelieu42/chimera/src/core/file/fileKit"
	"github.com/xuri/excelize/v2"
)

// NewEmptyFile 新建1个空白的Excel文件（其内只有1个空白的工作表，表名为"Sheet1"）
/*
@param path 文件的路径（如果文件已经存在，会覆盖它）
*/
func NewEmptyFile(path string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.MkParentDirs(path); err != nil {
		return nil, err
	}

	f := excelize.NewFile()
	err := f.SaveAs(path, opts...)
	if err != nil {
		defer func(f *excelize.File) {
			_ = f.Close()
		}(f)
		return nil, err
	}
	return f, nil
}
