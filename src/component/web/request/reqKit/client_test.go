package reqKit

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imroc/req/v3"
	"github.com/richelieu-yang/chimera/v3/src/log/zapKit"
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
	zapKit.Infof("response contenty: %s", string(data))
}

// TestNewClient1 测试retry count
func TestNewClient1(t *testing.T) {
	client := NewClient(WithDev(), WithTimeout(time.Second*10), WithRetryCount(0), WithRetryInterval(func(resp *req.Response, attempt int) time.Duration {
		zapKit.Debugf("attempt: %d", attempt)
		return time.Second
	}))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	zapKit.Info("-")
	data, err := client.Post("http://127.0.0.1:8001/test").SetContext(ctx).Do().ToBytes()
	zapKit.Info("=")
	if err != nil {
		panic(err)
	}
	zapKit.Infof("response contenty: %s", string(data))
}
