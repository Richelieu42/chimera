package redisKit

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
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
		err := client.XGroupCreateMkStream(context.TODO(), "stream:test", "gg1", "$")
		if err != nil {
			if IsConsumerGroupNameAlreadyExistError(err) {
				console.Warnf("consumer group already exists")
			} else {
				console.Fatalf("error: %T %s", err, err.Error())
			}
		} else {
			console.Info("OK")
		}

		id, err := client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: "ccccccc",
			Values: map[string]interface{}{
				"a": 1,
			},
		})
		if err != nil {
			console.Fatal(err.Error())
		}
		console.Infof("id: %s", id)
	}
}
