package excelKit

import "github.com/xuri/excelize/v2"

const (
	//// MaxRow row的取值范围: [0, 1048575]
	//MaxRow = 1048575
	//// MaxCol col的取值范围: [0, 16383]
	//MaxCol = 16383

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
