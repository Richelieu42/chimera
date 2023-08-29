package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/richelieu-yang/chimera/v2/src/copyKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonKit"
)

type Bean struct {
	Id int
}

func main() {
	b := &Bean{
		Id: 666,
	}
	src := map[string]interface{}{
		"b":   false,
		"tmp": b,
	}
	var dest map[string]interface{}

	dest, err := copyKit.DeepCopy(src)
	if err != nil {
		panic(err)
	}

	////dest = deepcopy.Copy(src).(map[string]interface{})
	//
	////dest = DeepCopy(src).(map[string]interface{})
	//
	//dest = map[string]interface{}{}
	//if err := copier.CopyWithOption(&dest, src, copier.Option{
	//	DeepCopy: true,
	//}); err != nil {
	//	panic(err)
	//}

	fmt.Println(src)
	fmt.Println(dest)

	src["b"] = true
	b.Id = 777

	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, src))
	fmt.Println(jsonKit.MarshalToStringWithAPI(jsoniter.ConfigCompatibleWithStandardLibrary, dest))
}
