package etcdKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
)

type (
	Config struct {
		Endpoints []string `json:"endpoints" yaml:"endpoints"`
	}
)

func (config *Config) Check() error {
	if err := interfaceKit.AssertNotNil(config, "config"); err != nil {
		return err
	}

	config.Endpoints = sliceKit.Uniq(config.Endpoints)
	config.Endpoints = sliceKit.RemoveEmpty(config.Endpoints, true)
	if sliceKit.IsEmpty(config.Endpoints) {
		return errorKit.Newf("config.Endpoints is empty")
	}

	return nil
}
