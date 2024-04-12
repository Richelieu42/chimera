package pushKit

import (
	"net/http"
)

type Listener interface {
	OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string)

	OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel)

	// OnMessage 收到 客户端 发来的消息.
	/*
		PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
	*/
	OnMessage(channel Channel, messageType int, data []byte)

	// BeforeClosedByBackend 由于后端主动关闭连接，而触发的关闭事件.
	BeforeClosedByBackend(channel Channel, closeInfo string)

	OnClose(channel Channel, closeInfo string, bsid, user, group string)
}
