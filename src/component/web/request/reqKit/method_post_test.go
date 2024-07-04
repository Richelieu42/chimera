package reqKit

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
	"time"
)

func TestPost(t *testing.T) {
	go func() {
		port := 8001

		engine := gin.Default()
		engine.POST("/test", func(ctx *gin.Context) {
			name := ctx.PostForm("name")
			age := ctx.PostForm("age")
			ctx.String(200, fmt.Sprintf("Hello %s(%s)", name, age))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second * 3)

	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	//resp := Post(context.TODO(), "http://127.0.0.1:8001/test", map[string]interface{}{
	//	"name": "张3",
	//	"age":  "30",
	//})
	resp := Post(context.TODO(), "http://127.0.0.1:8001/test", &person{
		Name: "张3",
		Age:  30,
	})
	if resp.Err != nil {
		panic(resp.Err)
	}
	fmt.Println(resp.ToString())
}
