package strKit

// SubBefore
/*
@param index 不包括
@return 范围: [0, index)
*/
func SubBefore(str string, index int) string {
	return str[:index]
}

// SubAfter
/*
@param index 包括
@return 范围: [index, length)
*/
func SubAfter(str string, index int) string {
	return str[index:]
}

// Substring 类似：Java的String.substring()
/*
@param from	开始的下标（包括）
@param to	结束的下标（不包括）
@return 范围: [from, to)

要点:
(1) from和to都不能 < 0
(2) 必须满足条件: from <= to（如果from == to，将返回""）
(3) 下标不能越界

e.g.
("abcd", 1, 1) => ""
*/
func Substring(str string, from, to int) string {
	return str[from:to]
}
