package wsKit

import (
	"github.com/gorilla/websocket"
	"github.com/richelieu-yang/chimera/v3/src/compress/gzipKit"
)

type MessageType struct {
	*gzipKit.Config

	value int
}

var (
	MessageTypeText = &MessageType{
		Config: nil,
		value:  websocket.TextMessage,
	}

	MessageTypeBinary = &MessageType{
		Config: nil,
		value:  websocket.BinaryMessage,
	}
)

// NewGzipMessageType
/*
PS: 此种情况下，必定使用 websocket.BinaryMessage.
*/
func NewGzipMessageType(level, compressThreshold int) (*MessageType, error) {
	if err := gzipKit.AssertValidLevel(level); err != nil {
		return nil, err
	}

	return &MessageType{
		Config: &gzipKit.Config{
			Level:             level,
			CompressThreshold: compressThreshold,
		},
		value: websocket.BinaryMessage,
	}, nil
}
