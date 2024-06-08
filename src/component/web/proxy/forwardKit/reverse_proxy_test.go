package forwardKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/logKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
)

/*
访问url: http://127.0.0.1/test
效果: 	将 http://127.0.0.1/test 转发给 http://127.0.0.1:8000/test
*/
func TestNewSingleHostReverseProxyWithUrl(t *testing.T) {
	port := 80

	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		errLog := logKit.NewStdoutLogger("")
		rp, err := NewSingleHostReverseProxyWithUrl("http://127.0.0.1:8000", errLog)
		if err != nil {
			ctx.String(500, err.Error())
			return
		}
		// (1) Richelieu: 此处应该为 true
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)

		if err := rp.Forward(ctx.Writer, ctx.Request); err != nil {
			// (2) Richelieu: 此处应该为 true
			fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
			ctx.String(500, err.Error())
			return
		}
		// (2) Richelieu: 此处应该为 true
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
		return
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
