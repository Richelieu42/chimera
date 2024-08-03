package mapKit

import "github.com/gogf/gf/v2/container/gmap"

/* map[interface{}]interface{} */
var (
	// NewTreeMap
	/*
	   PS: 当需要让返回结果按照自然升序排列时使用TreeMap.

	   @param comparator e.g. gutil.ComparatorInt
	*/
	NewTreeMap func(comparator func(v1, v2 interface{}) int, safe ...bool) *gmap.TreeMap = gmap.NewTreeMap

	NewTreeMapFrom func(comparator func(v1, v2 interface{}) int, data map[interface{}]interface{}, safe ...bool) *gmap.TreeMap = gmap.NewTreeMapFrom
)
