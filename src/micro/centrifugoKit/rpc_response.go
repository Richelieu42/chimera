package centrifugoKit

import (
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v3/src/micro/centrifugoKit/proxyproto"
)

// PackToRpcResponse
/*
@param base64Flag true: 对响应内容进行base64编码
*/
func PackToRpcResponse(jsonData []byte, base64Flag bool) *proxyproto.RPCResponse {
	resp := &proxyproto.RPCResponse{
		Result: &proxyproto.RPCResult{
			B64Data: "",
			Data:    nil,
		},
		Error:      nil,
		Disconnect: nil,
	}

	if base64Flag {
		resp.Result.B64Data = base64Kit.EncodeToString(jsonData)
	} else {
		resp.Result.Data = jsonData
	}
	return resp
}

// PackStringToRpcResponse
/*
@param base64Flag true: 对响应内容进行base64编码
*/
func PackStringToRpcResponse(jsonStr string, base64Flag bool) *proxyproto.RPCResponse {
	resp := &proxyproto.RPCResponse{
		Result: &proxyproto.RPCResult{
			B64Data: "",
			Data:    nil,
		},
		Error:      nil,
		Disconnect: nil,
	}

	if base64Flag {
		resp.Result.B64Data = base64Kit.EncodeStringToString(jsonStr)
	} else {
		resp.Result.Data = proxyproto.Raw(jsonStr)
	}
	return resp
}
