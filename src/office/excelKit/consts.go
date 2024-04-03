package excelKit

import "github.com/xuri/excelize/v2"

const (
	// MaxRows
	/*
		PS:
		(1) 从 1 开始
		(2) 范围: [1, 1048576]
	*/
	MaxRows = excelize.TotalRows

	// MaxColumns
	/*
		PS:
		(1) 从 1 开始
		(2) 范围: [1, 16384]
	*/
	MaxColumns = excelize.MaxColumns
)
