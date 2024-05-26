package centrifugoKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/proxyproto"
)

func ParseToRpcRequest(ctx *gin.Context) (*proxyproto.RPCRequest, error) {
	rpcRequest := &proxyproto.RPCRequest{}
	if err := ctx.Bind(rpcRequest); err != nil {
		return nil, err
	}
	return rpcRequest, nil
}

func PackToRpcResponse(json []byte) *proxyproto.RPCResponse {
	return &proxyproto.RPCResponse{
		Result: &proxyproto.RPCResult{
			B64Data: base64Kit.EncodeToString(json),
		},
		Error:      nil,
		Disconnect: nil,
	}
}

func PackStringToRpcResponse(jsonStr string) *proxyproto.RPCResponse {
	return &proxyproto.RPCResponse{
		Result: &proxyproto.RPCResult{
			B64Data: base64Kit.EncodeStringToString(jsonStr),
		},
		Error:      nil,
		Disconnect: nil,
	}
}
