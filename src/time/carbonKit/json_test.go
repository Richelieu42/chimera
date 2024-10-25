package carbonKit

//import (
//	"fmt"
//	"github.com/golang-module/carbon/v2"
//	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
//	"testing"
//)
//
///*
//参考: https://github.com/golang-module/carbon/blob/master/README.cn.md#json
//
//PS: 如果 carbon 标签没有设置，默认是 layout:2006-01-02 15:04:05；如果 tz 标签没有设置，默认是 Local.
//*/
//func TestJson(t *testing.T) {
//	type Person struct {
//		Name string `json:"name"`
//		Age  int    `json:"age"`
//
//		Birthday0 carbon.Carbon `json:"birthday0"`
//		Birthday1 carbon.Carbon `json:"birthday1" carbon:"layout:2006-01-02"`
//		Birthday2 carbon.Carbon `json:"birthday2" carbon:"layout:15:04:05"`
//		Birthday3 carbon.Carbon `json:"birthday3" carbon:"layout:2006-01-02 15:04:05"`
//		Birthday4 carbon.Carbon `json:"birthday4" carbon:"layout:2006-01-02" tz:"PRC"`
//		Birthday5 carbon.Carbon `json:"birthday5" carbon:"layout:15:04:05" tz:"PRC"`
//		Birthday6 carbon.Carbon `json:"birthday6" carbon:"layout:2006-01-02 15:04:05" tz:"PRC"`
//	}
//
//	now := Parse("2020-08-05 13:14:15", carbon.PRC)
//	person := Person{
//		Name: "gouguoyin",
//		Age:  18,
//
//		Birthday0: now,
//		Birthday1: now,
//		Birthday2: now,
//		Birthday3: now,
//		Birthday4: now,
//		Birthday5: now,
//		Birthday6: now,
//	}
//
//	loadErr := carbon.LoadTag(&person)
//	if loadErr != nil {
//		// 错误处理
//		panic(loadErr)
//	}
//
//	jsonStr, err := jsonKit.MarshalIndentToString(person, "", "   ")
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(jsonStr)
//}
