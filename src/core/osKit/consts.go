package osKit

import "os"

const (
	// PathSeparator 路径分隔符（string类型）
	/*
		e.g.
		"/": Mac、Linux
		"\": Windows
	*/
	PathSeparator = string(os.PathSeparator)

	// PathListSeparator 路径列表分隔符（string类型）
	/*
		e.g.
		":": Mac、Linux
		";": Windows
	*/
	PathListSeparator = string(os.PathListSeparator)
)
