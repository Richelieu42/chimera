package reqKit

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	go func() {
		port := 8001

		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(200, fmt.Sprintf("This is [%d].", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	// 等一会，让web服务先启动
	time.Sleep(time.Millisecond * 100)

	client := NewClient()
	//client := NewClient(WithDev())
	data, err := client.Post("http://127.0.0.1:8001/test").Do().ToBytes()
	if err != nil {
		panic(err)
	}
	console.Infof("response contenty: %s", string(data))
}

// TestNewClient1 测试retry count
func TestNewClient1(t *testing.T) {
	retryCount := 0

	client := NewClient(WithDev(), WithTimeout(time.Second*10))

	client.SetCommonRetryCount(retryCount)
	//client.SetCommonRetryFixedInterval(time.Millisecond * 100)
	client.SetCommonRetryInterval(func(resp *req.Response, attempt int) time.Duration {
		console.Debugf("attempt: %d", attempt)
		return time.Millisecond * 100
	})
	client.AddCommonRetryCondition(func(resp *req.Response, err error) bool {
		return err != nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	console.Info("-")
	data, err := client.Post("http://127.0.0.1:8001/test").SetContext(ctx).Do().ToBytes()
	console.Info("=")
	if err != nil {
		panic(err)
	}
	console.Infof("response contenty: %s", string(data))
}

func TestNewClient2(t *testing.T) {
	go func() {
		port := 8001

		engine := gin.Default()
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(500, fmt.Sprintf("This is [%d].", port))
		})
		if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
			panic(err)
		}
	}()

	time.Sleep(time.Second * 3)

	client := NewClient(WithDev())
	resp := client.Get("http://127.0.0.1:8001/test").Do()
	if resp.Err != nil {
		panic(resp.Err)
	}
}
