package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

// Round 四舍五入，保留n位小数.
/*
	@return 小数位可能小于传参n e.g. Round(3.1029, 2) => 3.1

   	e.g.
	fmt.Println(Round(3.14, 1))  // 3.1
   	fmt.Println(Round(3.15, 1))  // 3.2
   	fmt.Println(Round(-3.14, 1)) // -3.1
   	fmt.Println(Round(-3.15, 1)) // -3.2
*/
func Round[T constraints.Float | constraints.Integer](x T, n int) float64 {
	return mathutil.RoundToFloat(x, n)
}

// RoundToString 四舍五入，保留n位小数，返回字符串
func RoundToString[T constraints.Float | constraints.Integer](x T, n int) string {
	return mathutil.RoundToString(x, n)
}

// TruncRound 截断n位小数（不进行四舍五入）.
/*
   @param n 保留的小数位（可以 < 0，但有点奇怪!!!）

   e.g.
   (1234.124, 0)	=> 1234
   (1234.124, -1)	=> 1234
   (1234.124, -2)	=> 0

   (100.125, 2)	=> 100.12
   (100.125, 3)	=> 100.125
*/
func TruncRound[T constraints.Float | constraints.Integer](x T, n int) T {
	return mathutil.TruncRound(x, n)
}

//// Round 保留小数位（四舍五入），类似于 math.Round()，但功能更强大
///*
//PS:
//(1) 参考：https://zhuanlan.zhihu.com/p/152050239?from_voters_page=true
//(2) 个人感觉：先把正负号拿出来 => 进行取舍 => 把正负号还回去
//(3) 传参为负数的情况下，Golang的四舍五入与别的语言（Java、JavaScript）不同，详情见"Golang.docx"中的"math标准库".
//
//@param places 小数位数（如果最后几个都是0的话，会省略掉）；可以为负值
//
//e.g.
//(3.14, 1)	=> 3.1
//(3.15, 1)	=> 3.2
//(-3.14, 1) 	=> -3.1
//(-3.15, 1) 	=> -3.2
//
//(3.1001, 2) => 3.1
//(521, -1) 	=> 520
//*/
//func Round(f float64, places int32) float64 {
//	f, _ = decimal.NewFromFloat(f).Round(places).Float64()
//	return f
//}

// Ceil 向上取整，类似于 math.Ceil()，但功能更强大.
/*
PS:
(1) NOTE: this will panic on NaN, +/-inf
(2) 个人感觉: x轴向右.

e.g.
	(3.14, 1)	=> 3.2
	(-3.14, 1)	=> -3.1
*/
func Ceil(f float64, places int) float64 {
	f, _ = decimal.NewFromFloat(f).RoundCeil(int32(places)).Float64()
	return f
}

// Floor 向下取整，类似于 math.Floor()，但功能更强大.
/*
PS:
(1) NOTE: this will panic on NaN, +/-inf
(2) 个人感觉: x轴向左.

e.g.
	(3.14, 1)	=> 3.1
	(-3.14, 1)	=> -3.2
*/
func Floor(f float64, places int) float64 {
	f, _ = decimal.NewFromFloat(f).RoundFloor(int32(places)).Float64()
	return f
}
