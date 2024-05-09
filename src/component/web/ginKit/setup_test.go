package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
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
		engine.Any("/api.do", func(ctx *gin.Context) {
			base64Str := base64Kit.EncodeStringToString(`{"msg":"hello"}`)
			resp := &RPCResponse{
				Result: &RPCResult{
					B64Data: base64Str,
				},
			}
			jsonStr, err := jsonKit.MarshalToString(resp)
			if err != nil {
				panic(err)
			}
			ctx.String(200, jsonStr)
		})

		return nil
	}, WithServiceInfo("TEST"), WithDefaultFavicon(true))
}

type Raw []byte

type RPCResponse struct {
	Result *RPCResult `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

type RPCResult struct {
	Data    Raw    `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	B64Data string `protobuf:"bytes,2,opt,name=b64data,proto3" json:"b64data,omitempty"`
}
