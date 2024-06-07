package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	port := 80
	u, err := url.Parse("http://127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		rp := httputil.NewSingleHostReverseProxy(u)
		rp.ModifyResponse = func(response *http.Response) error {
			if response.StatusCode != http.StatusOK {
				return errors.New(fmt.Sprintf("invalid status code(%d)", response.StatusCode))
			}
			return nil
		}
		ctx.String(200, fmt.Sprintf("This is [%d].", port))
	})
	if err := engine.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}
