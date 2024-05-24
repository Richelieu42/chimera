package jsonKit

import "github.com/richelieu-yang/chimera/v3/src/core/strKit"

func IsJsonString(str string) (rst bool) {
	if strKit.StartWith(str, "{") {
		rst = strKit.EndWith(str, "}")
	} else if strKit.StartWith(str, "[") {
		rst = strKit.EndWith(str, "]")
	}
	return
}

func IsJson(data []byte) (rst bool) {
	length := len(data)
	if length >= 2 {
		if data[0] == '{' {
			rst = data[length-1] == '}'
		} else if data[0] == '[' {
			rst = data[length-1] == ']'
		}
	}
	return
}
