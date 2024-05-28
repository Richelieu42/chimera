package centrifugoKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/proxyproto"
)

// ParseToRpcRequest
/*
case 0: centrifugo的 proxy_binary_encoding 配置为 false（默认），请求参数取 RPCRequest.Data
case 1: centrifugo的 proxy_binary_encoding 配置为 true，请求参数取 RPCRequest.B64Data，还需要手动base64解码一下(with base64.StdEncoding)
*/
func ParseToRpcRequest(ctx *gin.Context) (*proxyproto.RPCRequest, error) {
	rpcRequest := &proxyproto.RPCRequest{}
	if err := ctx.Bind(rpcRequest); err != nil {
		return nil, err
	}
	return rpcRequest, nil
}
