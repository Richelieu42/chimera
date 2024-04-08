package sseKit

import (
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"time"
)

// NewProcessor
/*
!!!: 需要先调用 pushKit.MustSetUp 或 pushKit.SetUp.

@param idGenerator	可以为nil（将使用xid）
@param listener		不能为nil
@param msgType		消息类型
@param pongInterval	pong的周期（<=0则不发送pong）
*/
func NewProcessor(idGenerator func() (string, error), listener pushKit.Listener, msgType messageType, pongInterval time.Duration) (pushKit.Processor, error) {
	if err := pushKit.CheckSetup(); err != nil {
		return nil, err
	}

	if idGenerator == nil {
		idGenerator = pushKit.DefaultIdGenerator()
	}
	listeners, err := pushKit.NewListeners(listener, true)
	if err != nil {
		return nil, err
	}

	processor := &SseProcessor{
		idGenerator: idGenerator,
		listeners:   listeners,
		msgType:     msgType,
	}
	return processor, nil
}
