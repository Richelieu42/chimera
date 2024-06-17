// Package floatKit
package floatKit

import (
	"strconv"
)

// FormatFloat32ToString 类型转换: float32 => string
func FormatFloat32ToString(f float32, fmt byte, prec int) string {
	return strconv.FormatFloat(float64(f), fmt, prec, 32)
}

// FormatFloat64ToString 类型转换: float64 => string
/*
@param fmt	(1) "f": -ddd.dddd（十进制）
			(2) "b": -ddddp±ddd，指数为二进制
			(3) "e": -d.dddde±dd，十进制指数
			(4) "E": -d.ddddE±dd，十进制指数
			(5) "g": 指数很大时用"e"格式，否则"f"格式
			(6) "G": 指数很大时用"E"格式，否则"f"格式
@param prec (1) 如果传参fmt为"f"、"e"、"E"，它表示小数点后的数字个数
			(2) 如果传参fmt为"g"、"G"，它控制总的数字个数
			(3) -1: 使用"最少数量但又必需"的数字来表示传参f
*/
func FormatFloat64ToString(f float64, fmt byte, prec int) string {
	return strconv.FormatFloat(f, fmt, prec, 64)
}

// ToReadableString 去除后面的无意义的"0"
/*
PS:
(1) 十进制；
(2) 去掉后面的无意义的"0"；
(3) 同类的方法: humanize.Ftoa().

e.g.
(2.24)			=> "2.24"
(2.0000)		=> "2"
(2.000010000)	=> "2.00001"
*/
func ToReadableString(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}
