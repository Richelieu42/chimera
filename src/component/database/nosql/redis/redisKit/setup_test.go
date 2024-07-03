package redisKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
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
		fmt.Println("---")
		fmt.Println(client.Set(context.TODO(), "c", "zzz", 0))
		fmt.Println(client.Get(context.TODO(), "c"))
		fmt.Println("---")

		//fmt.Println(client.Ping(context.TODO()))

		//fmt.Println(client.IsStreamSupported(context.Background()))

		//m := map[string]interface{}{}
		//m["use"] = "pulsar"
		//m["pulsar"] = &pulsarKit.Config{
		//	Addrs: []string{"localhost:6650"},
		//}
		//
		//jsonStr, err := jsonKit.MarshalToString(m)
		//if err != nil {
		//	panic(err)
		//}
		//count, err := client.Publish(context.TODO(), "WsMq", jsonStr)
		//if err != nil {
		//	panic(err)
		//}
		//fmt.Println(count)
	}
}
