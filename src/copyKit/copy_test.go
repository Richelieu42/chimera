package copyKit

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
	"testing"
)

func TestDeepCopy(t *testing.T) {
	type bean struct {
		Id int
	}

	b := &bean{
		Id: 666,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": b,
	}

	/* (1) 浅拷贝 */
	//dest := &map[string]interface{}{}
	//err := Copy(dest, src)
	//if err != nil {
	//	panic(err)
	//}

	/* (2) 深拷贝 */
	dest := DeepCopy(src)

	///* (3) 深拷贝1 */
	//dest, err := DeepCopy1(src)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(src)
	fmt.Println(dest)

	// 修改src的内容（并不会影响dest）
	src["b"] = true
	b.Id = 777

	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, src))
	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, dest))
}

// case: 传参为 nil
func TestDeepCopy1(t *testing.T) {
	var m map[string]interface{} = nil

	m1 := DeepCopy(m)
	fmt.Println(m1)        // map[]
	fmt.Println(m1 == nil) // true
}
