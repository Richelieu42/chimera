package jsonKit

import (
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu42/chimera/v2/src/assert/strAssert"
	"github.com/richelieu42/chimera/v2/src/core/ptrKit"
	"github.com/richelieu42/chimera/v2/src/core/sliceKit"
)

// Unmarshal 反序列化.
/*
@param ptr 	(1) 不能为nil
			(2) 指针类型
@param data	必要条件: len(data) > 0（包含: 不能为nil）
*/
func Unmarshal(ptr interface{}, data []byte) error {
	/* 传参检查 */
	if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
		return err
	}
	if err := sliceKit.AssertNotEmpty(data); err != nil {
		return err
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
	if err := ptrKit.AssertNotNilAndIsPointer(ptr); err != nil {
		return err
	}
	if err := strAssert.AssertStringNotBlank(str); err != nil {
		return err
	}

	return jsoniter.UnmarshalFromString(str, ptr)
}
