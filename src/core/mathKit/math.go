package mathKit

import (
	"github.com/duke-git/lancet/v2/mathutil"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"golang.org/x/exp/constraints"
)

var (
	// Exponent 指数计算（x的n次方）.
	/*
		@param x 底数
		@param n 指数
		@return x^n

		e.g.
		(2, 10)	=> 1024
		(8, 2)	=> 64
	*/
	Exponent func(x, n int64) int64 = mathutil.Exponent

	// Factorial 阶乘.
	/*
		在数学中，阶乘是一种运算符号，表示所有小于等于该数的正整数的积12345。
		例如，5的阶乘表示为5!，其值为1×2×3×4×5=12012345。此外，0的阶乘被定义为112345。

		e.g.
		fmt.Println(mathKit.Factorial(0)) // 1
		fmt.Println(mathKit.Factorial(1)) // 1（=1）
		fmt.Println(mathKit.Factorial(5)) // 120（=1*2*3*4*5）
	*/
	Factorial func(x uint) uint = mathutil.Factorial

	// IsPrime 判断质数.
	IsPrime func(n int) bool = mathutil.IsPrime

	// Fibonacci 计算斐波那契数列的第n个数.
	Fibonacci func(first, second, n int) int = mathutil.Fibonacci

	// PointDistance 计算两个坐标点的距离.
	PointDistance func(x1, y1, x2, y2 float64) float64 = mathutil.PointDistance

	// AngleToRadian 将角度值转为弧度值.
	AngleToRadian func(angle float64) float64 = mathutil.AngleToRadian

	// RadianToAngle 将弧度值转为角度值.
	RadianToAngle func(radian float64) float64 = mathutil.RadianToAngle

	// Sin 正弦函数（计算弧度的正弦值）.
	Sin func(radian float64, precision ...int) float64 = mathutil.Sin

	// Cos 余弦函数（计算弧度的余弦值）.
	Cos func(radian float64, precision ...int) float64 = mathutil.Cos

	// Log 计算以base为底n的对数.
	Log func(n, base float64) float64 = mathutil.Log
)

// Average 计算平均数（可能需要对结果调用RoundToFloat方法四舍五入）
func Average[T constraints.Integer | constraints.Float](numbers ...T) float64 {
	return mathutil.Average[T](numbers...)
}

// Percent 计算百分比，保留 places 位小数
/*
@param places 保留的小数位
@return (val * 100 / total)

e.g.
	fmt.Println(mathKit.Percent(1, 3, 3)) // 33.333
*/
func Percent(val, total float64, places int32) float64 {
	//return mathutil.Percent(val, total, n)

	if val == 0 || total == 0 {
		return 0
	}
	d := decimal.NewFromFloat(val).Mul(decimal.NewFromInt32(100)).Div(decimal.NewFromFloat(total))
	f, _ := d.Round(places).Float64()
	return f
}

// GCD 计算最大公约数。
func GCD[T constraints.Integer](integers ...T) T {
	return mathutil.GCD(integers...)
}

// LCM 计算最小公倍数。
func LCM[T constraints.Integer](integers ...T) T {
	return mathutil.LCM(integers...)
}

// Clamp clamps number within the inclusive lower and upper bounds.
/*
case value < min: 	返回min
case value > max: 	返回max
case others:		返回value

e.g.
(0, -10, 10)	=> 0
(-42, -10, 10)	=> -10
(42, -10, 10)	=> 10
*/
func Clamp[T constraints.Ordered](value T, min T, max T) T {
	return lo.Clamp(value, min, max)
}

// Abs 绝对值.
func Abs[T constraints.Integer | constraints.Float](x T) T {
	return mathutil.Abs(x)
}
