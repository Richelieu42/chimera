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

	OnClose(channel Channel, closeInfo string, bsid, user, group string)
}

type Listeners []Listener

func (listeners Listeners) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	for _, listener := range listeners {
		listener.OnFailure(w, r, failureInfo)
	}
}

func (listeners Listeners) OnHandshake(w http.ResponseWriter, r *http.Request, channel Channel) {
	for _, listener := range listeners {
		listener.OnHandshake(w, r, channel)
	}
}

func (listeners Listeners) OnMessage(channel Channel, messageType int, data []byte) {
	for _, listener := range listeners {
		listener.OnMessage(channel, messageType, data)
	}
}

func (listeners Listeners) OnClose(channel Channel, closeInfo string) {
	// !!!: 此处先把取出来，以防 inner listener 解绑时去掉了相关信息（bsid, user, group），不会去掉id、data，导致轮到 另一个listener 时取不到数据
	bsid := channel.GetBsid()
	user := channel.GetUser()
	group := channel.GetGroup()

	for _, listener := range listeners {
		listener.OnClose(channel, closeInfo, bsid, user, group)
	}
}
