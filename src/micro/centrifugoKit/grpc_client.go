package centrifugoKit

import (
	"context"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
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

func NewGrpcClient(hosts []string) (client *GrpcClient, err error) {
	hosts = sliceKit.Uniq(hosts)
	hosts = sliceKit.RemoveEmpty(hosts, true)
	if err = sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return
	}

	// TODO
	logrus.Warn("TODO")
	client = &GrpcClient{}
	return
}
