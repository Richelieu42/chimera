package rocketmq5Kit

import (
	rmq_client "github.com/apache/rocketmq-clients/golang/v5"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/strKit"
)

// NewSimpleConsumer
/*
PS:
(1) In most case, you don't need to create many consumers, singletion pattern is more recommended.
(2) 需要先 set up!!!
(3) 第一个返回值，非nil的情况下，不再需要时请调用 GracefulStop().

@param consumerGroup 			不能为""
@param subscriptionExpressions	(1) key: 	topic，不能为 "*" || blank
								(2) value: 	tag，一般为 rmq_client.SUB_ALL
@param clientLogPath 			客户端日志（blank则输出到控制台）
*/
func NewSimpleConsumer(consumerGroup string, subscriptionExpressions map[string]*rmq_client.FilterExpression) (rmq_client.SimpleConsumer, error) {
	if config == nil {
		return nil, NotSetupError
	}

	if err := strKit.AssertNotEmpty(consumerGroup, "consumerGroup"); err != nil {
		return nil, err
	}
	if err := mapKit.AssertNotEmpty(subscriptionExpressions, "subscriptionExpressions"); err != nil {
		return nil, err
	}
	endpoint := sliceKit.Join(config.Endpoints, ";")

	simpleConsumer, err := rmq_client.NewSimpleConsumer(&rmq_client.Config{
		Endpoint:      endpoint,
		ConsumerGroup: consumerGroup,
		Credentials:   config.Credentials,
	},
		rmq_client.WithAwaitDuration(AwaitDuration),
		rmq_client.WithSubscriptionExpressions(subscriptionExpressions),
	)
	if err != nil {
		return nil, errorKit.Wrap(err, "Fail to new simple consumer")
	}
	if err := simpleConsumer.Start(); err != nil {
		return nil, errorKit.Wrap(err, "Fail to start simple consumer")
	}
	return simpleConsumer, nil
}
