package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/proxyKit"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		logrus.Infof("wd: [%s].", wd)
	}

	type config struct {
		Gin *Config `json:"gin" yaml:"gin"`
	}

	path := "_chimera-lib/config.yaml"
	c := &config{}
	//err := yamlKit.UnmarshalFromFile(path, c)
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Gin, func(engine *gin.Engine) error {
		target := "127.0.0.1:12000"

		engine.Any("/test", func(ctx *gin.Context) {
			if err := proxyKit.ProxyWithGin(ctx, target); err != nil {
				if !proxyKit.IsNegligibleError(err) {
					logrus.WithError(err).Error("Fail to proxy.")
					ctx.AbortWithStatus(500)
					return
				}
				logrus.WithError(err).Debug("A negligible error.")
				return
			}
			return
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true), WithDefaultNoRouteHtml(true))
}
