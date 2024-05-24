package centrifugoKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/apiproto"
	"google.golang.org/grpc"
)

type GrpcClient struct {
	mutexKit.Mutex
	apiproto.CentrifugoApiClient

	conn *grpc.ClientConn
}

func (client *GrpcClient) Close() (err error) {
	client.LockFunc(func() {
		if client.conn != nil {
			err = client.conn.Close()
			client.conn = nil
		}
	})
	return
}

func (client *GrpcClient) PublishSimply(ctx context.Context, channel string, data []byte) (*apiproto.PublishResponse, error) {
	in := &apiproto.PublishRequest{
		Channel: channel,
		Data:    apiproto.Raw(data),
	}
	return client.Publish(ctx, in)
}

func (client *GrpcClient) BroadcastSimply(ctx context.Context, channels []string, data []byte) (*apiproto.BroadcastResponse, error) {
	in := &apiproto.BroadcastRequest{
		Channels: channels,
		Data:     apiproto.Raw(data),
	}
	return client.Broadcast(ctx, in)
}
