package wsKit

import "github.com/gorilla/websocket"

type MessageType struct {
	value int
}

var (
	MessageTypeText = &MessageType{
		value: websocket.TextMessage,
	}

	MessageTypeBinary = &MessageType{
		value: websocket.BinaryMessage,
	}
)

func NewGzipMessageType() *MessageType {
	return &MessageType{
		value: websocket.BinaryMessage,
	}
}
