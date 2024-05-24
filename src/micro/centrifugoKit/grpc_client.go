package centrifugoKit

import (
	"context"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
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
@param hosts		centrifugo服务的grpc地址列表 e.g.[]string{"127.0.0.1:10000", "127.0.0.1:10001"}
@param scheme		grpc客户端负载均衡(slb)使用的scheme，可以为nil（将自动生成）
@param grpcApiKey	对应centrifugo服务配置文件中的 "grpc_api_key"
*/
func NewGrpcClient(hosts []string, scheme string, grpcApiKey string) (client *GrpcClient, err error) {
	hosts = sliceKit.PolyfillStringSlice(hosts)
	if err = sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return
	}
	if err = validateKit.Var(hosts, "required,gte=1,unique,dive,hostname_port"); err != nil {
		err = errorKit.Wrapf(err, "hosts is invalid")
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
