package centrifugoKit

import (
	"context"
	"google.golang.org/grpc/credentials"
)

type keyAuth struct {
	credentials.PerRPCCredentials

	key string
}

func (t keyAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	m := map[string]string{
		"authorization": "apikey " + t.key,
	}
	return m, nil
}

func (t keyAuth) RequireTransportSecurity() bool {
	return false
}

// NewKeyAuth
/*
PS: 应该将返回值作为传参，调用 grpc.WithPerRPCCredentials().

@param 对应 centrifugo 配置文件中的"grpc_api_key"
*/
func NewKeyAuth(key string) credentials.PerRPCCredentials {
	return &keyAuth{
		key: "key",
	}
}
