package redisKit

import (
	"github.com/richelieu42/chimera/v2/src/confKit"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSetUp(t *testing.T) {
	path := "/Users/richelieu/GolandProjects/chimera/chimera-lib/config.yaml"

	type config struct {
		Redis *Config `json:"redis"`
	}
	c := &config{}
	confKit.MustLoad(path, c)
	MustSetUp(c.Redis)

	client, err := GetClient()
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Info(client != nil)
}
