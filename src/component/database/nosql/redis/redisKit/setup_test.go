package redisKit

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
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

	console.Info("---")
	{
		if err := client.XGroupCreateMkStream(context.Background(), "stream:test", "group", "$"); err != nil {
			if !IsConsumerGroupNameAlreadyExistError(err) {
				console.Fatal(err.Error())
			}
		}

		for i := range 1000 {
			id, err := client.XAdd(context.Background(), &redis.XAddArgs{
				MaxLen: 100,
				Approx: true,

				Stream: "stream:test",
				Values: map[string]interface{}{
					"id": idKit.NewXid(),
				},
			})
			if err != nil {
				console.Errorf("i: %d id: %s", i, err.Error())
			} else {
				console.Infof("i: %d id: %s", i, id)
			}
		}

		//entries, err := client.XReadGroup(context.Background(), &redis.XReadGroupArgs{
		//	Streams:  []string{"stream:test", ">"},
		//	Group:    "group",
		//	Consumer: "consumer",
		//	Count:    10,
		//	//Block:    0,
		//	//NoAck:    false,
		//})
		//if err != nil {
		//	console.Error(err.Error())
		//} else {
		//	console.Infof("length: %d", len(entries))
		//}
	}
}
