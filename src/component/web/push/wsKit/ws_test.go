package wsKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

type TestListener struct {
	types.Listener
}

func (listener *TestListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("info", failureInfo).Info("OnFailure")
}

func (listener *TestListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel types.Channel) {
	logrus.Info("OnHandshake")

	_ = channel.Push([]byte("hello"))

	//// 一连接成功，后端就主动断开连接
	//_ = channel.Close()
}

// OnMessage 收到 客户端 发来的消息.
/*
	PS: 仅适用于WebSocket连接，因为SSE连接是单工的.
*/
func (listener *TestListener) OnMessage(channel types.Channel, messageType int, data []byte) {
	logrus.WithFields(logrus.Fields{
		"messageType": messageType,
		"text":        string(data),
	}).Info("OnMessage")
}

func (listener *TestListener) OnClose(channel types.Channel, closeInfo string) {
	logrus.WithField("info", closeInfo).Info("OnClose")
}

// TestNewProcessor
/*
url: ws://127.0.0.1/ws
*/
func TestNewProcessor(t *testing.T) {
	listener := &TestListener{}
	processor, err := NewProcessor(nil, nil, listener, MessageTypeBinary)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	engine.GET("/ws", func(ctx *gin.Context) {
		processor.Process(ctx.Writer, ctx.Request)
	})
	if err := engine.Run(":80"); err != nil {
		panic(err)
	}
}
