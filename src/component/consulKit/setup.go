package consulKit

import (
	"github.com/hashicorp/consul/api"
	"github.com/richelieu-yang/chimera/v3/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var innerClient *api.Client

func MustSetUp(config *api.Config) {
	if err := SetUp(config); err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(config *api.Config) (err error) {
	innerClient, err = NewClient(config)
	return
}
