//go:build !go1.22

package randomKit

import "github.com/duke-git/lancet/v2/random"

var (
	// RandFloat 生成随机float64数字，可以指定范围和精度.
	/*
		@param precision 精度（小数点后保留几位）
		@return [min, max)
	*/
	RandFloat func(min, max float64, precision int) float64 = random.RandFloat

	// RandFloatSlice 生成随机float64数字切片，指定长度，范围和精度.
	/*
		@param precision 精度（小数点后保留几位）
		@return (1) 切片内的元素范围: [min, max)
				(2) 切片内的元素不会重复
	*/
	RandFloatSlice func(n int, min, max float64, precision int) []float64 = random.RandFloats
)
