package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"net/http"
	"time"
)

// DefaultUpgrader 默认的Upgrader.
/*
@return 并发安全的
*/
func DefaultUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		HandshakeTimeout: time.Second * 3,
		CheckOrigin: func(r *http.Request) bool {
			// 允许跨域
			return true
		},
	}
}

// NewProcessor
/*
!!!: 需要先调用 pushKit.MustSetUp 或 pushKit.SetUp.

@param upgrader			可以为nil（将使用默认的）
@param idGenerator		可以为nil（将使用xid）
@param listener			不能为nil
*/
func NewProcessor(upgrader *websocket.Upgrader, idGenerator func() (string, error), listener pushKit.Listener, messageType messageType) (*WsProcessor, error) {
	if err := pushKit.CheckSetup(); err != nil {
		return nil, err
	}

	if upgrader == nil {
		upgrader = DefaultUpgrader()
	}
	if idGenerator == nil {
		idGenerator = pushKit.DefaultIdGenerator()
	}
	listeners, err := pushKit.NewListeners(listener, false)
	if err != nil {
		return nil, err
	}

	return &WsProcessor{
		upgrader:    upgrader,
		idGenerator: idGenerator,
		listeners:   listeners,
		msgType:     messageType,
	}, nil
}
