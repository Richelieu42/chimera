package sseKit

import (
	"encoding/base64"
	"fmt"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v3/src/crypto/base64Kit"
	"github.com/richelieu-yang/chimera/v3/src/time/timeKit"
	"github.com/richelieu-yang/chimera/v3/src/urlKit"
	"net/http"
	"time"
)

type SseChannel struct {
	pushKit.BaseChannel

	w        http.ResponseWriter
	r        *http.Request
	msgType  messageType
	interval *timeKit.Interval
}

func (channel *SseChannel) Initialize() error {
	channel.interval = timeKit.NewInterval(func(t time.Time) {
		if err := channel.Push(pushKit.PongData); err != nil {
			pushKit.GetDefaultLogger().WithError(err).Error("Fail to pong")
			return
		}
	}, channel.PongInterval)
	return nil
}

// Dispose 仅是释放资源，不会关闭通道（应当先关闭通道，再释放资源）.
func (channel *SseChannel) Dispose() {
	channel.interval.Stop()
	channel.interval = nil
}

// Push （写锁）推送消息给客户端.
func (channel *SseChannel) Push(data []byte) error {
	return channel.PushMessage(channel.msgType, data)
}

// PushMessage （写锁）推送消息给客户端.
func (channel *SseChannel) PushMessage(msgType messageType, data []byte) (err error) {
	var str string
	switch msgType {
	case MessageTypeEncode:
		str = string(data)
		str = urlKit.EncodeURIComponent(str)
	case MessageTypeBase64:
		str = base64Kit.EncodeToString(data, base64Kit.WithEncoding(base64.StdEncoding))
	case MessageTypeRaw:
		fallthrough
	default:
		str = string(data)
	}
	event := &Event{
		Data: str,
	}

	if channel.Closed {
		return pushKit.ChannelClosedError
	}
	/* 写锁 */
	channel.LockFunc(func() {
		if channel.Closed {
			err = pushKit.ChannelClosedError
			return
		}
		err = event.Push(channel.w)
	})
	return err
}

// Close （写锁）后端主动关闭通道.
func (channel *SseChannel) Close(reason string) error {
	closeInfo := fmt.Sprintf("Closed by backend with reason(%s)", reason)

	channel.Listeners.BeforeClosedByBackend(channel, closeInfo)

	if channel.SetClosed() {
		channel.CloseCh <- closeInfo
	}
	return nil
}

func (channel *SseChannel) BindGroup(group string) {
	pushKit.BindGroup(channel, group)
}

func (channel *SseChannel) BindUser(user string) {
	pushKit.BindUser(channel, user)
}

func (channel *SseChannel) BindBsid(bsid string) {
	pushKit.BindBsid(channel, bsid)
}
