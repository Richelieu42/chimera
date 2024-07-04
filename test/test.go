package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"time"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	go func() {
		port := 8001

		engine := gin.Default()
		engine.POST("/test", func(ctx *gin.Context) {
			p := &person{}
			if err := ctx.Bind(p); err != nil {
				ctx.String(500, err.Error())
				return
			}
			ctx.String(200, fmt.Sprintf("Hello %s(%d)", p.Name, p.Age))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second)

	// 创建请求客户端
	client := req.C()

	// 设置请求数据
	p := &person{
		Name: "李四",
		Age:  40,
	}

	// 发送POST请求
	resp, err := client.R().
		//SetContentType("application/json; charset=utf-8").
		SetBody(p).
		Post("http://127.0.0.1:8001/test")

	if err != nil {
		panic(err)
	}

	// 处理响应
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", resp.String())
}
