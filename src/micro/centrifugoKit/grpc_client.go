package centrifugoKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/consts"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/sliceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/apiproto"
	"github.com/richelieu-yang/chimera/v3/src/validateKit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewGrpcClient centrifugo服务的grpc客户端，支持客户端负载均衡（slb）
/*
@param hosts		centrifugo服务的grpc地址列表 e.g.[]string{"127.0.0.1:10000", "127.0.0.1:10001"}
@param scheme		grpc客户端负载均衡(slb)使用的scheme
					(1) 可以为""，将自动生成
					(2) 其中不能有大写字母
					(3) 可以有: 小写字母、数字、"-"...
@param grpcApiKey	对应centrifugo服务配置文件中的 "grpc_api_key"
*/
func NewGrpcClient(hosts []string, scheme string, grpcApiKey string) (*GrpcClient, error) {
	/* hosts */
	hosts = sliceKit.PolyfillStringSlice(hosts)
	if err := sliceKit.AssertNotEmpty(hosts, "hosts"); err != nil {
		return nil, err
	}
	if err := validateKit.Var(hosts, "required,gte=1,unique,dive,hostname_port"); err != nil {
		err = errorKit.Wrapf(err, "hosts is invalid")
		return nil, err
	}

	/* scheme */
	if strKit.IsEmpty(scheme) {
		scheme = fmt.Sprintf("%s-centrifugo-grpc-client-slb-%s", consts.LowerProjectName, idKit.NewXid())
	} else {
		scheme = strKit.ToLower(scheme)
	}

	/* grpcApiKey */
	if err := strKit.AssertNotEmpty(grpcApiKey, "grpcApiKey"); err != nil {
		return nil, err
	}

	// Richelieu: target中的"hello"随意，甚至可以去掉
	target := fmt.Sprintf("%s:///hello", scheme)

	conn, err := grpc.NewClient(target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithPerRPCCredentials(NewKeyAuth(grpcApiKey)),
	)
	if err != nil {
		return nil, err
	}
	client := apiproto.NewCentrifugoApiClient(conn)
	return &GrpcClient{
		CentrifugoApiClient: client,
		conn:                conn,
	}, nil
}
