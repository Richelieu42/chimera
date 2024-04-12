package wsKit

import "github.com/gorilla/websocket"

type messageType struct {
	value int
}

var (
	MessageTypeText = messageType{
		value: websocket.TextMessage,
	}

	MessageTypeBinary = messageType{
		value: websocket.BinaryMessage,
	}
)
