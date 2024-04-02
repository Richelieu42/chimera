package pulsarKit

import (
	"context"
	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
)

type (
	Consumer struct {
		pulsar.Client
		pulsar.Consumer
	}
)

func (c *Consumer) Close() {
	if c.Consumer != nil {
		c.Consumer.Close()
	}
	if c.Client != nil {
		c.Client.Close()
	}
}

// NewConsumerOriginally
/*
PS: 目标Pulsar服务未启动的情况下，如果ctx不加以限制，要过约 1min 才会返回error（期间客户端日志有connection refused输出）.

@param options			必须的属性: Topic、SubscriptionName、Type
@param clientLogPath 	客户端的日志输出（为空则输出到控制台; 不会rotate）
*/
func NewConsumerOriginally(ctx context.Context, addresses []string, options pulsar.ConsumerOptions, clientLogPath string) (*Consumer, error) {
	var client pulsar.Client
	var consumer pulsar.Consumer

	// 写入nil: 新建Consumer成功
	errCh := make(chan error, 1)

	go func() {
		var err error
		client, err = NewClient(addresses, clientLogPath)
		if err != nil {
			errCh <- err
			return
		}
		consumer, err = client.Subscribe(options)
		if err != nil {
			err = errorKit.Wrapf(err, "client fails to subscribe with topic(%s), subscriptionName(%s) and type(%s)",
				options.Topic, options.SubscriptionName, options.Type)
			errCh <- err
			return
		}

		select {
		case <-ctx.Done():
			// 新建Consumer成功之前，ctx已经 超时 或 被取消 了，此时需要释放资源
			if consumer != nil {
				consumer.Close()
			}
			if client != nil {
				client.Close()
			}
			errCh <- ctx.Err()
		default:
			errCh <- nil
		}
	}()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case err := <-errCh:
		if err != nil {
			return nil, err
		}
		return &Consumer{
			Client:   client,
			Consumer: consumer,
		}, nil
	}
}
