package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/vulcand/oxy/v2/buffer"
	"github.com/vulcand/oxy/v2/forward"
	"github.com/vulcand/oxy/v2/roundrobin"
	"net/url"
)

func main() {
	// Forwards incoming requests to whatever location URL points to, adds proper forwarding headers
	//fwd := forward.New(false)
	/*
		@param passHostHeader 是否传递请求头中的Host头？
		(1) true:	被代理服务收到请求，请求头中包含Host: 127.0.0.1:8000（原始请求的Host）
		(2) false:	被代理服务收到请求，请求头中包含Host: localhost:8001（Host被修改了，现在是被代理服务的Host）
		(3) 建议使用: true
	*/
	fwd := forward.New(false)
	lb, err := roundrobin.New(fwd)
	if err != nil {
		panic(err)
	}

	servers := []string{
		"http://localhost:8001",
		"http://localhost:8002",
	}
	for _, s := range servers {
		u, err := url.Parse(s)
		if err != nil {
			panic(err)
		}
		if err := lb.UpsertServer(u); err != nil {
			panic(err)
		}
	}

	// buf will read the request body and will replay the request again in case if forward returned status
	// corresponding to nework error (e.g. Gateway Timeout)
	buf, err := buffer.New(lb, buffer.Retry(`IsNetworkError() && Attempts() < 2`))
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		logrus.Infof("[%s]", ctx.Request.Header.Get("Host"))
		logrus.Infof("Host: [%s]", ctx.Request.Host)
		logrus.Infof("URL.Host: [%s]", ctx.Request.URL.Host)

		buf.ServeHTTP(ctx.Writer, ctx.Request)
	})
	if err := engine.Run(":8000"); err != nil {
		panic(err)
	}
}
