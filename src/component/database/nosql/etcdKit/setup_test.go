package etcdKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/config/viperKit"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	{
		wd, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName)
		if err != nil {
			panic(err)
		}
		console.Infof("working dir: %s", wd)
	}

	type config struct {
		Etcd *Config `json:"etcd"`
	}

	path := "_chimera-lib/config.yaml"
	c := &config{}
	//err := yamlKit.UnmarshalFromFile(path, c)
	_, err := viperKit.UnmarshalFromFile(path, nil, c)
	if err != nil {
		panic(err)
	}

	MustSetUp(c.Etcd, "")
	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
	}
	kv := clientv3.NewKV(client)

	resp, err := kv.Put(context.TODO(), "k", "v")
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(resp.Header.GetRevision())
}
