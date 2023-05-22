package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/ptrKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

// Unmarshal 反序列化.
/*
@param ptr 	(1) 不能为nil
			(2) 指针类型
@param data	必要条件: len(data) > 0（包含: 不能为nil）
*/
func Unmarshal(ptr interface{}, data []byte) error {
	/* 传参检查 */
	if ptr == nil {
		return errorKit.Simple("ptr == nil")
	}
	if !ptrKit.IsPointer(ptr) {
		return errorKit.Simple("type(%T) of ptr isn't pointer", ptr)
	}
	if len(data) == 0 {
		if data == nil {
			return errorKit.Simple("data == nil")
		}
		return errorKit.Simple("len(data) == 0")
	}

	return jsoniter.Unmarshal(data, ptr)
}

// UnmarshalFromString 反序列化.
/*
@param ptr 	(1) 不能为nil
			(2) 指针类型
@param str	不能为空字符串("")
*/
func UnmarshalFromString(ptr interface{}, str string) error {
	/* 传参检查 */
	if ptr == nil {
		return errorKit.Simple("ptr == nil")
	}
	if !ptrKit.IsPointer(ptr) {
		return errorKit.Simple("type(%T) of ptr isn't pointer", ptr)
	}
	if strKit.IsEmpty(str) {
		return errorKit.Simple("str is empty")
	}

	return jsoniter.UnmarshalFromString(str, ptr)
}
