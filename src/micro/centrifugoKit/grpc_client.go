package centrifugoKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/sirupsen/logrus"
)

type GrpcClient struct {
}

func (client *GrpcClient) Batch(ctx context.Context) error {
	// TODO
	logrus.Warn("TODO")

	return nil
}

func (client *GrpcClient) Publish(ctx context.Context, channel string, data []byte) error {
	// TODO
	logrus.Warn("TODO")

	return nil
}

func (client *GrpcClient) Broadcast(ctx context.Context, channels []string, data []byte) error {
	// TODO
	logrus.Warn("TODO")

	return nil
}

func (client *GrpcClient) Presence(ctx context.Context, channel string) error {
	// TODO
	logrus.Warn("TODO")

	return nil
}

func (client *GrpcClient) PresenceStats(ctx context.Context, channel string) error {
	// TODO
	logrus.Warn("TODO")

	return nil
}

// NewGrpcClient
/*

 */
func NewGrpcClient(hosts []string, grpcApiKey string) (client *GrpcClient, err error) {
	hosts = sliceKit.RemoveEmpty(hosts, true)
	hosts = sliceKit.Uniq(hosts)
	if err = sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return
	}

	scheme := fmt.Sprintf("%s-centrifugo-grpc-client-slb-%s", consts.LowerProjectName, idKit.NewXid())
	//// Richelieu: scheme里面不能有大写字母
	//scheme = strKit.ToLower(scheme)
	target := fmt.Sprintf("%s:///hello", scheme)

	// TODO
	logrus.Warn("TODO")
	client = &GrpcClient{}
	return
}
