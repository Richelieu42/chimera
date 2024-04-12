package wsKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v3/src/compress/gzipKit"
	"time"
)

var (
	_ pushKit.Channel = &WsChannel{}
)

type WsChannel struct {
	pushKit.BaseChannel

	msgType *MessageType

	conn *websocket.Conn
}

func (channel *WsChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage 推送消息给客户端.
/*
@param MessageType MessageTypeText || MessageTypeBinary
*/
func (channel *WsChannel) PushMessage(messageType *MessageType, data []byte) (err error) {
	/* gzip压缩 */
	if messageType.gzipConfig != nil {
		data, err = gzipKit.CompressComplexly(data, gzipKit.WithLevel(messageType.gzipConfig.Level), gzipKit.WithCompressThreshold(messageType.gzipConfig.CompressThreshold))
		if err != nil {
			return err
		}
	}

	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	abortFlag := false

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			abortFlag = true
			return
		}

		err = channel.conn.WriteMessage(messageType.value, data)
	})

	if abortFlag {
		return
	}
	// Closed == true 的情况下，推送消息失败（基本上就是连接断开了）
	closeInfo := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
	if channel.SetClosed() {
		channel.CloseCh <- closeInfo
	}
	return
}

// Close 后端主动关闭通道.
func (channel *WsChannel) Close(reason string) (err error) {
	closeInfo := fmt.Sprintf("Closed by backend with reason(%s)", reason)

	channel.Listeners.BeforeClosedByBackend(channel, closeInfo)

	// Richelieu: 这里延迟3毫秒，以免: BeforeClosedByBackend()里面会向前端发消息，前端先触发onclose，再触发onmessage
	time.Sleep(time.Millisecond * 3)

	if channel.SetClosed() {
		channel.CloseCh <- closeInfo
		err = channel.conn.Close()
	}
	return
}
