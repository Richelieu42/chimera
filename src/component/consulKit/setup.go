package consulKit

import (
	"github.com/hashicorp/consul/api"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/log/console"
)

var (
	NotSetupError = errorKit.Newf("haven’t been set up correctly")
)

var innerClient *api.Client

func MustSetUp(config *api.Config) {
	if err := SetUp(config); err != nil {
		console.Fatalf("Fail to set up, error: %s", err.Error())
	}
}

func SetUp(config *api.Config) (err error) {
	innerClient, err = NewClient(config)
	return
}
