package redisKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/mq/pulsarKit"
	"github.com/richelieu-yang/chimera/v2/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v2/src/serialize/json/jsonKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetUp(t *testing.T) {
	wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
	if err != nil {
		panic(err)
	}
	logrus.Infof("wd: [%s].", wd)

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
		logrus.Fatal(err)
	}
	client = client

	{
		//fmt.Println(client.Ping(context.TODO()))

		m := map[string]interface{}{}
		m["use"] = "pulsar"
		m["pulsar"] = &pulsarKit.Config{
			Addrs: []string{"localhost:6650"},
		}

		jsonStr, err := jsonKit.MarshalToString(m)
		if err != nil {
			panic(err)
		}
		count, err := client.Publish(context.TODO(), "WsMq", jsonStr)
		if err != nil {
			panic(err)
		}
		fmt.Println(count)
	}
}
