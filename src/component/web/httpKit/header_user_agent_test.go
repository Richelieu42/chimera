package httpKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
)

// 访问url: http://127.0.0.1/test
func TestGetUserAgentInfo(t *testing.T) {
	port := 80

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		info := GetUserAgentInfo(ctx.Request.Header)
		fmt.Println("Browser", info.GetBrowser())
		fmt.Println("OS", info.GetOS())
		fmt.Println("Version", info.GetVersion())
		fmt.Println("IsDesktop", info.IsDesktop())
		fmt.Println("IsMobile", info.IsMobile())
		fmt.Println("IsTablet", info.IsTablet())
		fmt.Println("IsTV", info.IsTV())
		fmt.Println("IsBot", info.IsBot())

		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
