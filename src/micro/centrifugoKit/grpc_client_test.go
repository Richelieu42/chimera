package centrifugoKit

import (
	"context"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNewGrpcClient(t *testing.T) {
	hosts := []string{"172.18.21.50:13001", "172.18.21.50:13002"}
	scheme := "ccc"
	grpcApiKey := "2098e3e6-a41d-4004-8101-dbc3229e4cee"

	client, err := NewGrpcClient(hosts, scheme, grpcApiKey)
	if err != nil {
		logrus.Fatal(err)
	}
	defer client.Close()

	{
		tmp, err := client.PublishSimply(context.TODO(), "test", []byte("Publish"))
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(tmp)
	}

	{
		tmp, err := client.BroadcastSimply(context.TODO(), []string{"test"}, []byte("Broadcast"))
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.Info(tmp)
	}
}
