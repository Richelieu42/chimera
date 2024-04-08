package wsKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v3/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v3/src/concurrency/poolKit"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
	"time"
)

type demoListener struct {
	pushKit.Listener
}

func (l *demoListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("failureInfo", failureInfo).Error("OnFailure")
}

func (l *demoListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),
	}).Info("OnHandshake")
}

func (l *demoListener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"messageType": messageType,
		"text":        string(data),
	}).Info("OnMessage")

	text := fmt.Sprintf("length of received message: %d", len(data))
	if err := channel.Push([]byte(text)); err != nil {
		logrus.WithError(err).Error("Fail to push when on message.")
		return
	}
}

func (l *demoListener) BeforeClosedByBackend(channel pushKit.Channel, closeInfo string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"closeInfo": closeInfo,
	}).Info("BeforeClosedByBackend")
}

func (l *demoListener) OnClose(channel pushKit.Channel, closeInfo string, bsid, user, group string) {
	logrus.WithFields(logrus.Fields{
		"clientIP": channel.GetClientIP(),
		"type":     channel.GetType(),
		"id":       channel.GetId(),

		"closeInfo": closeInfo,
	}).Info("OnClose")
}

func TestWs(t *testing.T) {
	engine := gin.Default()
	engine.Use(ginKit.NewCorsMiddleware(nil))

	/* 初始化poolKit */
	pool, err := poolKit.NewAntPool(1000)
	if err != nil {
		panic(err)
	}
	pushKit.MustSetUp(pool, nil)
	//pushKit.MustSetUp(pool, nil, pushKit.WithWsPongInterval(-1))

	/* WebSocket */
	processor, err := NewProcessor(nil, nil, &demoListener{}, MessageTypeText, time.Second*3)
	if err != nil {
		logrus.Fatal(err)
	}
	engine.GET("/ws", processor.ProcessWithGin)

	if err := engine.Run(":12000"); err != nil {
		logrus.Fatal(err)
	}
}
