//go:build go1.22

package randomKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"math/rand/v2"
)

// RandFloat 生成随机float64数字，可以指定范围和精度.（参考: random.RandFloat）
/*
	@param precision 精度（小数点后保留几位）
	@return [min, max)
*/
func RandFloat(min, max float64, precision int) float64 {
	if min == max {
		return min
	}

	if max < min {
		min, max = max, min
	}

	n := rand.Float64()*(max-min) + min

	return mathutil.RoundToFloat(n, precision)
}

// RandFloatSlice 生成随机float64数字切片，指定长度，范围和精度.（参考: random.RandFloats）
/*
	@param precision 精度（小数点后保留几位）
	@return (1) 切片内的元素范围: [min, max)
			(2) 切片内的元素不会重复
*/
func RandFloatSlice(n int, min, max float64, precision int) []float64 {
	nums := make([]float64, n)
	used := make(map[float64]struct{}, n)
	for i := 0; i < n; {
		r := RandFloat(min, max, precision)
		if _, use := used[r]; use {
			continue
		}
		used[r] = struct{}{}
		nums[i] = r
		i++
	}

	return nums
}
