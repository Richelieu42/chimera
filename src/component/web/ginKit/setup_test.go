package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"github.com/sirupsen/logrus"
	"strings"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
		if err := pathKit.SetTempDir("_temp"); err != nil {
			panic(err)
		}

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
			//type bean struct {
			//	Name string `json:"name"`
			//}
			//
			//b := &bean{}
			////if err := ctx.ShouldBindJSON(b); err != nil {
			//if err := ctx.ShouldBind(b); err != nil {
			//	ctx.String(500, err.Error())
			//	return
			//}
			//name := b.Name

			//name := ObtainPostParam(ctx, "name")
			ctx.String(200, fmt.Sprintf("[%s] Hello %s.", timeKit.FormatCurrent(""), strings.Repeat("c", 200)))
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}
