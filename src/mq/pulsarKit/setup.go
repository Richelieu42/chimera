package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v2/src/log/logrusKit"
	"github.com/sirupsen/logrus"
)

var config *Config

func MustSetUp(config Config) {
	err := SetUp(config)
	if err != nil {
		logrusKit.DisableQuote(nil)
		logrus.Fatalf("%+v", err)
	}
}

func SetUp(pulsarConfig Config) (err error) {
	defer func() {
		if err != nil {
			config = nil
		}
	}()

	config = &pulsarConfig
	err = verify(config.VerifyConfig)
	return
}

// NewProducer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SendTimeout
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewProducer(ctx context.Context, options pulsar.ProducerOptions, clientLogPath string) (*Producer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewProducerOriginally(ctx, config.Addresses, options, clientLogPath)
}

// NewConsumer
/*
前提: 成功调用 SetUp() || MustSetUp().

@param options 必需属性: Topic、SubscriptionName、Type
@param logPath 客户端的日志输出（"": 输出到控制台）
*/
func NewConsumer(ctx context.Context, options pulsar.ConsumerOptions, clientLogPath string) (*Consumer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	return NewConsumerOriginally(ctx, config.Addresses, options, clientLogPath)
}
