package reqKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
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

	client := NewClient(WithDev())

	data, err := client.Post("http://127.0.0.1:8001/test").Do().ToBytes()
	if err != nil {
		panic(err)
	}
	fmt.Println("response:", string(data))
}
