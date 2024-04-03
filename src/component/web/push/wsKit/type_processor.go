package wsKit

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v3/src/component/web/httpKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/mutexKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"net/http"
	"time"
)

type WsProcessor struct {
	pushKit.Processor

	// upgrader 是并发安全的
	upgrader *websocket.Upgrader

	idGenerator func() (string, error)
	listeners   pushKit.Listeners
	msgType     messageType
}

func (p *WsProcessor) ProcessWithGin(ctx *gin.Context) {
	p.Process(ctx.Writer, ctx.Request)
}

func (p *WsProcessor) Process(w http.ResponseWriter, r *http.Request) {
	PolyfillWebSocketRequest(r)

	// 先判断是不是websocket请求
	if !IsWebSocketUpgrade(r) {
		failureInfo := "Not a websocket upgrade request"
		p.listeners.OnFailure(w, r, failureInfo)
		return
	}

	// Upgrade（升级为WebSocket协议）
	conn, err := p.upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		err = errorKit.Wrapf(err, "Fail to upgrade")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	// PS: 对于 Conn.Close() ，可以多次调用，不会panic，但从第二次关闭开始，返回非nil的error（可以直接忽略）.
	defer conn.Close()

	channel, err := p.newChannel(r, conn, make(chan string, 1))
	if err != nil {
		err = errorKit.Wrapf(err, "Fail to new channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}
	if err := channel.Initialize(); err != nil {
		err = errorKit.Wrapf(err, "Fail to initialize channel")
		p.listeners.OnFailure(w, r, err.Error())
		return
	}

	p.listeners.OnHandshake(w, r, channel)

	conn.SetCloseHandler(func(code int, text string) error {
		closeInfo := fmt.Sprintf("code: %d, text: %s", code, text)
		if channel.SetClosed() {
			channel.GetCloseCh() <- closeInfo
		}

		// 默认的close handler
		message := websocket.FormatCloseMessage(code, text)
		_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

	go func() {
		/* 接收WebSocket客户端发来的消息 */
		for {
			messageType, data, err := conn.ReadMessage()
			if err != nil {
				var closeInfo string
				var closeErr *websocket.CloseError
				if errors.As(err, &closeErr) {
					closeInfo = fmt.Sprintf("code: %d, text: %s", closeErr.Code, closeErr.Text)
				} else {
					closeInfo = fmt.Sprintf("Fail to read message because of error(%s)", err.Error())
				}

				if channel.SetClosed() {
					channel.GetCloseCh() <- closeInfo
				}
				break
			}
			p.listeners.OnMessage(channel, messageType, data)
		}
	}()

	select {
	case <-r.Context().Done():
		if channel.SetClosed() {
			p.listeners.OnClose(channel, "Context done")
		}
	case closeInfo := <-channel.GetCloseCh():
		p.listeners.OnClose(channel, closeInfo)
	}
}

func (p *WsProcessor) newChannel(r *http.Request, conn *websocket.Conn, closeCh chan string) (pushKit.Channel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, errorKit.Wrapf(err, "Fail to generate id with idGenerator")
	}
	if err := strKit.AssertNotBlank(id, "id"); err != nil {
		return nil, err
	}

	ip := httpKit.GetClientIP(r)
	channel := &WsChannel{
		BaseChannel: pushKit.BaseChannel{
			RWMutex:      mutexKit.RWMutex{},
			CloseCh:      closeCh,
			ClientIP:     ip,
			Type:         "WebSocket",
			Id:           id,
			Bsid:         "",
			User:         "",
			Group:        "",
			Data:         nil,
			Closed:       false,
			Listeners:    p.listeners,
			PongInterval: pushKit.GetPongInterval(),
		},
		conn:        conn,
		messageType: p.msgType,
	}
	return channel, nil
}
