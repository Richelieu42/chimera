package melodyKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestNew(t *testing.T) {
	// 创建Melody实例
	m := melody.New()
	// 处理WebSocket连接
	m.HandleConnect(func(s *melody.Session) {
		fmt.Println("WebSocket connected")
	})
	// 处理WebSocket消息
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		logrus.Infof("Received a message: %s", string(msg))

		// 发送消息给所有连接的客户端
		text := "Broadcast " + string(msg)
		if err := m.Broadcast([]byte(text)); err != nil {
			logrus.WithError(err).Error("Fail to broadcast.")
		}
	})
	// 处理WebSocket断开连接
	m.HandleDisconnect(func(s *melody.Session) {
		fmt.Println("WebSocket disconnected")
	})

	engine := gin.Default()
	// 注册WebSocket处理函数
	engine.GET("/ws", func(c *gin.Context) {
		err := m.HandleRequest(c.Writer, c.Request)
		if err != nil {
			fmt.Println(err)
		}
	})
	if err := engine.Run(":8080"); err != nil {
		logrus.Fatal(err)
	}
}
