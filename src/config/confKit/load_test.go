package confKit

//
//import (
//	"fmt"
//	"testing"
//	"time"
//)
//
//func TestMustLoad(t *testing.T) {
//	type config1 struct {
//		Name string `json:"name,default=test"`
//	}
//
//	type config struct {
//		Time    time.Duration `json:"time"`
//		Number  int           `json:"number,range=[1:100)"`
//		Number1 int           `json:"number1,range=[1:)"`
//		// 此时使用 "config1" 而非 "*config1"，避免nil的情况
//		Config1 config1 `json:"config1"`
//	}
//
//	c := &config{}
//	MustLoad("./_test.yaml", c)
//
//	fmt.Printf("c.Time: %s\n", c.Time)
//	fmt.Printf("c.Number: %d\n", c.Number)
//	fmt.Printf("c.Number1: %d\n", c.Number1)
//	fmt.Printf("c.Config1.Name: %s\n", c.Config1.Name)
//}
