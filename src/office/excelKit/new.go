package excelKit

import (
	"github.com/richelieu-yang/chimera/v2/src/core/fileKit"
	"github.com/xuri/excelize/v2"
)

// NewFile
/*
PS: 返回的 *excelize.File实例，不再使用时应当调用 File.Close().
*/
var NewFile func(opts ...excelize.Options) *excelize.File = excelize.NewFile

// NewFileWithPath 新建1个空白的Excel文件（其内只有1个空白的工作表，表名为"Sheet1"）
/*
PS: 返回的 *excelize.File实例，不再使用时应当调用 File.Close().

@param path 文件的路径（如果文件已经存在，会覆盖它）
*/
func NewFileWithPath(path string, opts ...excelize.Options) (*excelize.File, error) {
	if err := fileKit.MkParentDirs(path); err != nil {
		return nil, err
	}

	f := NewFile()
	if err := f.SaveAs(path, opts...); err != nil {
		_ = f.Close()
		return nil, err
	}
	return f, nil
}
