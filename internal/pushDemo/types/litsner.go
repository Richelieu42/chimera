package types

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/sirupsen/logrus"
	"net/http"
)

type DemoListener struct {
	pushKit.Listener
}

func (l *DemoListener) OnFailure(w http.ResponseWriter, r *http.Request, failureInfo string) {
	logrus.WithField("failureInfo", failureInfo).Error("OnFailure")
}

func (l *DemoListener) OnHandshake(w http.ResponseWriter, r *http.Request, channel pushKit.Channel) {
	logrus.Info("OnHandshake")

	text := fmt.Sprintf("Hello, your id is [%s].", channel.GetId())
	if err := channel.Push([]byte(text)); err != nil {
		logrus.Error(err)
	}

	//go func() {
	//	time.Sleep(time.Second * 3)
	//	_ = channel.Close("测试")
	//}()
}

func (l *DemoListener) OnMessage(channel pushKit.Channel, messageType int, data []byte) {
	logrus.WithFields(logrus.Fields{
		"messageType": messageType,
		"text":        string(data),
	}).Info("OnMessage")
}

func (l *DemoListener) OnClose(channel pushKit.Channel, closeInfo string) {
	logrus.WithField("closeInfo", closeInfo).Info("OnClose")
}
