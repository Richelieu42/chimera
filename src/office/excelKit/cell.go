package excelKit

import "github.com/xuri/excelize/v2"

var (
	// CellNameToCoordinates 单元格名 => 坐标（行、列）
	/*
		e.g.
	*/
	CellNameToCoordinates func(cell string) (col int, row int, err error) = excelize.CellNameToCoordinates

	// CoordinatesToCellName 坐标（行、列） => 单元格名
	/*
		e.g.
	*/
	CoordinatesToCellName func(col, row int, abs ...bool) (cellName string, err error) = excelize.CoordinatesToCellName
)
