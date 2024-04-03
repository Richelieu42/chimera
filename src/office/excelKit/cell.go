package excelKit

import "github.com/xuri/excelize/v2"

var (
	// CellNameToCoordinates 单元格名 => 坐标（行、列）
	/*
		e.g.
			fmt.Println(excelKit.CellNameToCoordinates("A1")) // 1 1 <nil>
			fmt.Println(excelKit.CellNameToCoordinates("Z3")) // 26 3 <nil>
	*/
	CellNameToCoordinates func(cell string) (col int, row int, err error) = excelize.CellNameToCoordinates

	// CoordinatesToCellName 坐标（行、列） => 单元格名
	/*
		e.g.
			fmt.Println(excelKit.CoordinatesToCellName(1, 1))       // A1 <nil>
			fmt.Println(excelKit.CoordinatesToCellName(1, 1, true)) // $A$1 <nil>
	*/
	CoordinatesToCellName func(col, row int, abs ...bool) (cellName string, err error) = excelize.CoordinatesToCellName

	// JoinCellName
	/*
		e.g.
	*/
	JoinCellName func(col string, row int) (cellName string, err error) = excelize.JoinCellName

	// SplitCellName
	/*
		e.g.
	*/
	SplitCellName func(cell string) (col string, row int, err error) = excelize.SplitCellName
)
