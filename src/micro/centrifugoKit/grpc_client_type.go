package centrifugoKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/apiproto"
	"github.com/richelieu-yang/chimera/v3/src/serialize/json/jsonKit"
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

// PublishSimply
/*
@param data 必须是json!!!
*/
func (client *GrpcClient) PublishSimply(ctx context.Context, channel string, jsonData []byte) (*apiproto.PublishResponse, error) {
	if err := jsonKit.AssertJson(jsonData, "jsonData"); err != nil {
		return nil, err
	}

	in := &apiproto.PublishRequest{
		Channel: channel,
		Data:    apiproto.Raw(jsonData),
	}
	return client.Publish(ctx, in)
}

// BroadcastSimply
/*
@param data 必须是json!!!
*/
func (client *GrpcClient) BroadcastSimply(ctx context.Context, channels []string, jsonData []byte) (*apiproto.BroadcastResponse, error) {
	if err := jsonKit.AssertJson(jsonData, "jsonData"); err != nil {
		return nil, err
	}

	in := &apiproto.BroadcastRequest{
		Channels: channels,
		Data:     apiproto.Raw(jsonData),
	}
	return client.Broadcast(ctx, in)
}
