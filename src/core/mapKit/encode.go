package mapKit

import (
	"github.com/fatih/structs"
)

// Encode 结构体 => map[string]interface{}
/*
golang gorm 零值更新不生效问题处理
	https://blog.csdn.net/damanchen/article/details/124717553

PS: 一般情况下，需要使用 tag.

@param obj 	(1) 结构体（实例||指针）
			(2) 不能为nil（否则会panic）
@return 必定不为nil
*/
func Encode(obj interface{}) map[string]interface{} {
	s := structs.New(obj)
	// 不使用默认tag("structs")
	s.TagName = "json"
	return s.Map()
}

func EncodeWithTag(obj interface{}, tag string) map[string]interface{} {
	s := structs.New(obj)
	// 不使用默认tag("structs")
	s.TagName = tag
	return s.Map()
}
