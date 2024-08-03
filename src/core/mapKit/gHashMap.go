package mapKit

import "github.com/gogf/gf/v2/container/gmap"

/* map[interface{}]interface{} */
var (
	NewAnyAnyMap func(safe ...bool) *gmap.AnyAnyMap = gmap.NewAnyAnyMap

	NewAnyAnyMapFrom func(data map[interface{}]interface{}, safe ...bool) *gmap.AnyAnyMap = gmap.NewAnyAnyMapFrom
)

/* map[int]interface{} */
var (
	NewIntAnyMap func(safe ...bool) *gmap.IntAnyMap = gmap.NewIntAnyMap

	NewIntAnyMapFrom func(data map[int]interface{}, safe ...bool) *gmap.IntAnyMap = gmap.NewIntAnyMapFrom
)

/* map[string]interface{} */
var (
	NewStrAnyMap func(safe ...bool) *gmap.StrAnyMap = gmap.NewStrAnyMap

	NewStrAnyMapFrom func(data map[string]interface{}, safe ...bool) *gmap.StrAnyMap = gmap.NewStrAnyMapFrom
)

/* map[int]int */
var (
	NewIntIntMap func(safe ...bool) *gmap.IntIntMap = gmap.NewIntIntMap

	NewIntIntMapFrom func(data map[int]int, safe ...bool) *gmap.IntIntMap = gmap.NewIntIntMapFrom
)

/* map[string]string */
var (
	NewStrStrMap func(safe ...bool) *gmap.StrStrMap = gmap.NewStrStrMap

	NewStrStrMapFrom func(data map[string]string, safe ...bool) *gmap.StrStrMap = gmap.NewStrStrMapFrom
)

/* map[int]string */
var (
	NewIntStrMap func(safe ...bool) *gmap.IntStrMap = gmap.NewIntStrMap

	NewIntStrMapFrom func(data map[int]string, safe ...bool) *gmap.IntStrMap = gmap.NewIntStrMapFrom
)

/* map[string]int */
var (
	NewStrIntMap func(safe ...bool) *gmap.StrIntMap = gmap.NewStrIntMap

	NewStrIntMapFrom func(data map[string]int, safe ...bool) *gmap.StrIntMap = gmap.NewStrIntMapFrom
)
