package componentKit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/mq/rocketmq5Kit"
)

// NewProducer
/*
@param logConfig 可以为nil
*/
func NewProducer(logConfig *rocketmq5Kit.LogConfig) (rmq_client.Producer, error) {
	config, err := GetRocketmq5Config()
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errorKit.Simple("config == nil")
	}

	return rocketmq5Kit.NewProducer(logConfig, config)
}

// NewSimpleConsumer
/*
@param logConfig 可以为nil
*/
func NewSimpleConsumer(logConfig *rocketmq5Kit.LogConfig, consumerGroup, topic, tag string) (rmq_client.SimpleConsumer, error) {
	config, err := GetRocketmq5Config()
	if err != nil {
		return nil, err
	}
	if config == nil {
		return nil, errorKit.Simple("config == nil")
	}

	return rocketmq5Kit.NewSimpleConsumer(logConfig, config, consumerGroup, topic, tag)
}
