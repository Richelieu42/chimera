package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/apache/rocketmq-clients/golang/credentials"
	"github.com/richelieu42/chimera/v2/src/copyKit"
	"github.com/richelieu42/chimera/v2/src/core/errorKit"
	"github.com/richelieu42/chimera/v2/src/core/strKit"
)

type (
	Config struct {
		rmq_client.Config
		TopicToVerify string
	}
)

var (
	defaultCredentials = &credentials.SessionCredentials{
		AccessKey:    "",
		AccessSecret: "",
	}
)

// processConfig Consumer和Producer通用
func processConfig(baseConfig *rmq_client.Config) (*rmq_client.Config, error) {
	if baseConfig == nil {
		return nil, errorKit.New("config == nil")
	}

	// 深拷贝，为了不修改传参baseConfig
	var config *rmq_client.Config
	obj := copyKit.DeepCopy(baseConfig)
	config, ok := obj.(*rmq_client.Config)
	if !ok {
		return nil, errorKit.New("fail to deepcopy baseConfig")
	}

	if strKit.IsEmpty(config.Endpoint) {
		return nil, errorKit.New("config.Endpoint is empty")
	}

	config.ConsumerGroup = ""

	if config.Credentials == nil {
		config.Credentials = defaultCredentials
	}

	return config, nil
}
