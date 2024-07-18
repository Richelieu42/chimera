package redisKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
	"testing"
)

func TestSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		fmt.Println("wd:", wd)
	}

	type config struct {
		Redis *Config `json:"redis"`
	}
	path := "_chimera-lib/config.yaml"
	c := &config{}
	if _, err := viperKit.UnmarshalFromFile(path, nil, c); err != nil {
		panic(err)
	}
	fmt.Println(jsonKit.MarshalIndentToString(c.Redis, "", "    "))

	MustSetUp(c.Redis)
	client, err := GetClient()
	if err != nil {
		panic(err)
	}
	client = client

	{
		err := client.XGroupCreate(context.TODO(), "stream:test", "gg", "$")
		if err != nil {
			console.Error(err.Error())
		} else {
			console.Info("OK")
		}

		err = client.XGroupCreate(context.TODO(), "stream:test", "gg", "$")
		if err != nil {
			console.Error(err.Error())
		} else {
			console.Info("OK")
		}
	}
}
