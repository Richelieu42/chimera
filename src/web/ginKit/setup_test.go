package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/richelieu42/chimera/v2/src/dataSizeKit"
	"github.com/richelieu42/chimera/v2/src/log/logrusKit"
	"github.com/richelieu42/chimera/v2/src/web/httpKit"
	"net"
	"net/http"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	logrusKit.MustSetUp(nil)

	type config struct {
		Gin *Config `json:"Gin"`
	}
	c := &config{}

	confKit.MustLoad("/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml", c)
	MustSetUp(c.Gin, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			if err := httpKit.Proxy(ctx.Writer, ctx.Request, "http", "127.0.0.1:8001"); err != nil {
				if ee, ok := err.(*net.OpError); ok {
					eee := ee.Unwrap()
					eee = eee
					ctx.String(http.StatusInternalServerError, err.Error())
				} else {
					ctx.String(http.StatusInternalServerError, err.Error())
				}
			}
		})

		engine.Any("/upload", func(ctx *gin.Context) {
			fileHeader, err := ctx.FormFile("file")
			if err != nil {
				ctx.String(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.String(http.StatusOK, dataSizeKit.ToReadableStringWithIEC(uint64(fileHeader.Size)))
		})

		return nil
	})
}

func TestMustSetUp1(t *testing.T) {
	config := &Config{
		Port:       80,
		Colorful:   true,
		Middleware: nil,
		SSL: &SslConfig{
			CertFile: "/Users/richelieu/GolandProjects/chimera/chimera-lib/ssl.pem",
			KeyFile:  "/Users/richelieu/GolandProjects/chimera/chimera-lib/ssl.key",
			Port:     443,
		},
	}
	MustSetUp(config, nil, func(engine *gin.Engine) error {
		engine.Any("/test", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "ok")
		})
		return nil
	})
}
