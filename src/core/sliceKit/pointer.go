package sliceKit

// StringSlice 参考: playwright.StringSlice
func StringSlice(v ...string) *[]string {
	var o []string
	o = append(o, v...)
	return &o
}

func BoolSlice(v ...bool) *[]bool {
	var o []bool
	o = append(o, v...)
	return &o
}

// IntSlice 参考: playwright.IntSlice
func IntSlice(v ...int) *[]int {
	var o []int
	o = append(o, v...)
	return &o
}

func Int64Slice(v ...int64) *[]int64 {
	var o []int64
	o = append(o, v...)
	return &o
}
