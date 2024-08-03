package mapKit

import "github.com/gogf/gf/v2/container/gmap"

/* map[interface{}]interface{} */
var (
	// NewListMap
	/*
	   使用场景: 当需要按输入顺序返回结果时使用ListMap.
	*/
	NewListMap func(safe ...bool) *gmap.ListMap = gmap.NewListMap

	NewListMapFrom func(data map[interface{}]interface{}, safe ...bool) *gmap.ListMap = gmap.NewListMapFrom
)
