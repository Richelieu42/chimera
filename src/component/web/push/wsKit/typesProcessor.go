package wsKit

import (
	"errors"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"net/http"
	"sync"
	"time"
)

type Processor struct {
	// upgrader 是并发安全的
	upgrader *websocket.Upgrader

	idGenerator func() (string, error)

	listener pushKit.Listener
}

func (p *Processor) NewChannel(conn *websocket.Conn) (*WsChannel, error) {
	id, err := p.idGenerator()
	if err != nil {
		return nil, err
	}

	return &WsChannel{
		BaseChannel: pushKit.BaseChannel{
			Id:     id,
			Bsid:   "",
			User:   "",
			Group:  "",
			Mutex:  sync.Mutex{},
			Data:   nil,
			Closed: false,
		},
		conn: conn,
	}, nil
}

func (p *Processor) Handle(w http.ResponseWriter, r *http.Request) {
	PolyfillWebSocketRequest(r)

	// 先判断是不是websocket请求
	if !websocket.IsWebSocketUpgrade(r) {
		p.listener.OnFailure(w, r, "Not a websocket upgrade request")
		return
	}

	// Upgrade（升级为WebSocket协议）
	conn, err := upgrader.Upgrade(w, r, w.Header())
	if err != nil {
		err = errorKit.Wrap(err, "Fail to upgrade websocket")
		p.listener.OnFailure(w, r, err.Error())
		return
	}
	defer conn.Close()

	channel, err := p.NewChannel(conn)
	if err != nil {
		err = errorKit.Wrap(err, "Fail to new channel")
		p.listener.OnFailure(w, r, err.Error())
		return
	}

	p.listener.OnHandshake(w, r, channel)

	conn.SetCloseHandler(func(code int, text string) error {
		p.listener.OnClose(channel, code, text)

		//channel.SetClosed()
		//
		//if RemoveChannel(channel) {
		//	channel.GetListener().OnCloseByFrontend(channel, code, text)
		//}

		// 默认的close handler
		message := websocket.FormatCloseMessage(code, text)
		_ = conn.WriteControl(websocket.CloseMessage, message, time.Now().Add(time.Second))
		return nil
	})

	/* 接收WebSocket客户端发来的消息 */
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			channel.SetClosed()

			if RemoveChannel(channel) {
				var closeErr *websocket.CloseError
				if errors.As(err, &closeErr) {
					listener.OnCloseByFrontend(channel, closeErr.Code, closeErr.Text)
				} else {
					listener.OnCloseByBackend(channel)
				}
			}
			break
		}
		p.listener.OnMessage(channel, messageType, data)
	}
}

// NewProcessor
/*
@param handshakeTimeout	默认3s
@param checkOrigin		可以为nil（允许所有）
@param idGenerator		可以为nil（使用xid）
@param listener			不能为nil
*/
func NewProcessor(handshakeTimeout time.Duration, checkOrigin func(r *http.Request) bool, idGenerator func() (string, error), listener pushKit.Listener) (*Processor, error) {
	handshakeTimeout = timeKit.ToDefaultDurationIfInvalid(handshakeTimeout, time.Second*3)
	if checkOrigin == nil {
		checkOrigin = func(r *http.Request) bool {
			// 允许跨域
			return true
		}
	}
	upgrader := &websocket.Upgrader{
		HandshakeTimeout: handshakeTimeout,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}

	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}

	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return &Processor{
		upgrader:    upgrader,
		idGenerator: idGenerator,
		listener:    listener,
	}, nil
}
