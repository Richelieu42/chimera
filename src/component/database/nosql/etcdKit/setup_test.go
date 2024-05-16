package etcdKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/config/confKit"
	"github.com/richelieu-yang/chimera/v3/src/core/osKit"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	var path string
	if osKit.IsWindows() {
		path = "D:\\GolandProjects\\chimera\\chimera-lib\\config.yaml"
	} else {
		path = "/Users/richelieu/GolandProjects/chimera/_chimera-lib/config.yaml"
	}

	type config struct {
		Etcd *Config `json:"etcd"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
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
