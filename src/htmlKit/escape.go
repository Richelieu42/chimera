package htmlKit

import "html"

var (
	// EscapeString 转义html字符串
	/*
		PS: JavaScript可以使用 lodash 的 _.escape().
	*/
	EscapeString = html.EscapeString

	// UnescapeString 反转义html字符串
	UnescapeString = html.UnescapeString
)
