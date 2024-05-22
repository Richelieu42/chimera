package ginKit

import (
	"github.com/mholt/archiver/v4"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
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
		engine.Any("/test", func(ctx *gin.Context) {
			c := ctx.Request.Context()

			archiver.FilesFromDisk()

			select {
			case <-c.Done():
				logrus.Warn("c.Done()")
			case <-time.After(time.Second * 10):
				ctx.String(200, "OK")
			}

			//ctx.String(200, "ok")
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}
