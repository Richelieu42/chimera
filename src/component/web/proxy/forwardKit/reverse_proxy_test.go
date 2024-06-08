package forwardKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/log/logKit"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"testing"
)

/*
访问url: http://127.0.0.1:80/test
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
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)

		if err := rp.Forward(ctx.Writer, ctx.Request); err != nil {
			fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
			ctx.String(500, err.Error())
			return
		}
		fmt.Printf("rp.ErrorHandler == nil? [%t]\n", rp.ErrorHandler == nil)
		return
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}
