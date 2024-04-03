package wsKit

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"time"
)

type WsChannel struct {
	pushKit.BaseChannel

	conn        *websocket.Conn
	messageType messageType
	interval    *timeKit.Interval
}

func (channel *WsChannel) Initialize() error {
	channel.interval = timeKit.NewInterval(func(t time.Time) {
		if err := channel.Push(pushKit.PongData); err != nil {
			pushKit.GetDefaultLogger().WithError(err).Error("Fail to pong")
			return
		}
	}, channel.PongInterval)
	return nil
}

// Dispose 仅是释放资源，不会关闭通道（应当先关闭通道，再释放资源）.
func (channel *WsChannel) Dispose() {
	channel.interval.Stop()
	channel.interval = nil
}

func (channel *WsChannel) Push(data []byte) error {
	return channel.PushMessage(channel.messageType, data)
}

// PushMessage 推送消息给客户端.
/*
@param messageType MessageTypeText || MessageTypeBinary
*/
func (channel *WsChannel) PushMessage(messageType messageType, data []byte) (err error) {
	switch messageType {
	case MessageTypeText:
	case MessageTypeBinary:
	default:
		err = errorKit.Newf("invalid message type(%d)", messageType)
		return
	}
	if channel.Closed {
		return pushKit.ChannelClosedError
	}

	// 是否推送失败？
	failFlag := false

	/* 写锁 */
	channel.RWMutex.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}

		err = channel.conn.WriteMessage(messageType.value, data)
		failFlag = err != nil
	})

	if failFlag {
		// 推送消息失败，基本上就是连接断开了
		closeInfo := fmt.Sprintf("Fail to push because of error(%s)", err.Error())
		if channel.SetClosed() {
			channel.CloseCh <- closeInfo
		}
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

func (channel *WsChannel) BindGroup(group string) {
	pushKit.BindGroup(channel, group)
}

func (channel *WsChannel) BindUser(user string) {
	pushKit.BindUser(channel, user)
}

func (channel *WsChannel) BindBsid(bsid string) {
	pushKit.BindBsid(channel, bsid)
}
