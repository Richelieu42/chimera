package interfaceKit

import "reflect"

// IsNil Deprecated: 反射有性能问题，应尽可能避免使用此方法.
/*
go里面的类型包含 type 和 value 的，一般对于业务我们更在意value是不是空.

e.g.
	var src interface{} = nil
	var src1 []string = nil
	var src2 map[string]interface{} = nil
	type bean struct {
	}
	var src3 *bean = nil

	fmt.Println(interfaceKit.IsNil(src))  // true
	fmt.Println(interfaceKit.IsNil(src1)) // true
	fmt.Println(interfaceKit.IsNil(src2)) // true
	fmt.Println(interfaceKit.IsNil(src3)) // true
*/
func IsNil(obj interface{}) bool {
	//v := reflect.ValueOf(obj)
	//if v.Kind() == reflect.Ptr {
	//	// obj的类型：指针
	//	return v.IsNil()
	//}
	//// obj的类型：非指针
	//return obj == nil

	// golang中nil的判断 https://blog.csdn.net/weixin_44579563/article/details/129583860
	return obj == nil || reflect.ValueOf(obj).IsNil()
}
