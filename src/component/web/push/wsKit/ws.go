package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v2/src/core/timeKit"
	"github.com/richelieu-yang/chimera/v2/src/idKit"
	"net/http"
	"time"
)

// NewProcessor
/*
@param handshakeTimeout	默认3s
@param checkOrigin		可以为nil（允许所有）
@param idGenerator		可以为nil（使用xid）
@param listener			不能为nil
*/
func NewProcessor(handshakeTimeout time.Duration, checkOrigin func(r *http.Request) bool, idGenerator func() (string, error), listener pushKit.Listener) (*WsProcessor, error) {
	handshakeTimeout = timeKit.ToDefaultDurationIfInvalid(handshakeTimeout, time.Second*3)
	if checkOrigin == nil {
		checkOrigin = func(r *http.Request) bool {
			// 允许跨域
			return true
		}
	}
	upgrader := &websocket.Upgrader{
		HandshakeTimeout: handshakeTimeout,
		CheckOrigin:      checkOrigin,
	}
	if idGenerator == nil {
		idGenerator = func() (string, error) {
			return idKit.NewXid(), nil
		}
	}

	if err := interfaceKit.AssertNotNil(listener, "listener"); err != nil {
		return nil, err
	}

	return &WsProcessor{
		upgrader:    upgrader,
		idGenerator: idGenerator,
		listener:    listener,
	}, nil
}
