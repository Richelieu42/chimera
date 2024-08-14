package strKit

import "github.com/samber/lo"

// Substring Return part of a string.
/*
@param offset 开始下标（甚至可以<0, -1: 最后一个字符）

e.g.
	("hello", -1, 1)			=> "o"
	("hello", 2, 3)				=> "llo"
	("hello", -4, 3)			=> "ell"
	("hello", -2, math.MaxUint)	=> "lo"
*/
func Substring[T ~string](str T, offset int, length uint) T {
	return lo.Substring(str, offset, length)
}

//// Substring 类似：Java的String.substring()
///*
//@param from	开始的下标（包括）
//@param to	结束的下标（不包括）
//@return 范围: [from, to)
//
//要点:
//(1) from和to都不能 < 0
//(2) 必须满足条件: from <= to（如果from == to，将返回""）
//(3) 下标不能越界
//
//e.g.
//("abcd", 1, 1) => ""
//*/
//func Substring(str string, from, to int) string {
//	return str[from:to]
//}

// SubBefore
/*
@param index 不包括
@return 范围: [0, index)
*/
func SubBefore(s string, index int) string {
	return s[:index]
}

// SubBeforeString
/*
case 1: s包含str的情况，返回截取后的字符串；
case 2: s不包含str的情况，直接返回s.

e.g.
("abcd", "bc") => "a"
*/
func SubBeforeString(s, str string) string {
	i := Index(s, str)

	if i != -1 {
		return SubBefore(s, i)
	}
	return s
}

// SubAfter
/*
@param index 包括
@return 范围: [index, length)
*/
func SubAfter(s string, index int) string {
	return s[index:]
}

// SubAfterString
/*
case 1: s包含str的情况，返回截取后的字符串；
case 2: s不包含str的情况，直接返回s.

e.g.
("abcd", "bc") => "bcd"
*/
func SubAfterString(s, str string) string {
	i := Index(s, str)

	if i != -1 {
		return SubAfter(s, i)
	}
	return s
}
