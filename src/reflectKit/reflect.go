package reflectKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"reflect"
	"unsafe"
)

// GetField
/*
PS:
这样获取到的值是不可以修改的，如果要修改name字段，需要用到reflect.NewAt函数，这个函数通过一个类型值的底层地址（指针 p）和类型，
返回指向该值的一个指针，这个返回值是可寻址的，即可通过它直接访问该值。具体的修改是使用reflect.Value.Set函数。

@param fieldName 字段名（属性名），private 或 public的都可以
*/
func GetField(ptr interface{}, fieldName string) reflect.Value {
	return reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
}

// GetNestedField 获取（层层嵌套的）字段
func GetNestedField(ptr interface{}, fieldNames ...string) (reflect.Value, error) {
	if fieldNames == nil {
		return reflect.Value{}, errorKit.New("fieldNames == nil")
	}

	v := reflect.ValueOf(ptr).Elem()
	for i, name := range fieldNames {
		if i > 0 {
			/*
				参考: 	https://www.codenong.com/50098624/
				以避免: 	panic: reflect: call of reflect.Value.FieldByName on ptr Value
			*/
			v = reflect.Indirect(v)
		}
		v = v.FieldByName(name)
	}
	return v, nil
}

// SetField 设置字段的值.
/*
PS: 支持设置 未导出字段 的值.

@param fieldName 		字段名（属性名），private或public的都可以
@param newFieldValue	字段的新值
*/
func SetField(ptr interface{}, fieldName string, newFieldValue interface{}) error {
	v := GetField(ptr, fieldName)
	// 获取字段的可寻址反射对象（可set）
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()

	v1 := reflect.ValueOf(newFieldValue)
	if v.Kind() != v1.Kind() {
		return errorKit.New("expected kind %v, got kind: %v", v.Kind(), v1.Kind())
	}
	v.Set(v1)
	return nil
}
