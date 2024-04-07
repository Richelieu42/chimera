package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"github.com/sirupsen/logrus"
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
			//r := ctx.Request
			//readCloser, err := r.GetBody()
			//if err != nil {
			//	ctx.String(500, err.Error())
			//	return
			//}
			//readCloser = readCloser

			ctx.String(200, timeKit.FormatCurrent())

			ctx.String(200, `{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}{"errorMessage":"[O(192.168.80.115:9000)-0] no error","result":{"wdStaticPages":1},"errorCode":"0","appVersion":0}`)

			//qm := map[string][]string{
			//	"b": {"bOx"},
			//	"c": {"阿德去外地"},
			//}
			//
			//if err := proxyKit.ProxyWithGin(ctx, "127.0.0.1:10000", proxyKit.WithExtraQueryParams(qm)); err != nil {
			//	ctx.String(http.StatusInternalServerError, err.Error())
			//	return
			//}
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}
