package sseKit

type messageType struct {
	value string
}

var (
	// MessageTypeRaw 对于data，不做任何处理
	MessageTypeRaw = messageType{
		"raw",
	}

	// MessageTypeEncode 对于data，编码一下（前端需对应处理）
	MessageTypeEncode = messageType{
		"encode",
	}

	// MessageTypeBase64 对于data，base64编码一下（前端需对应处理）
	MessageTypeBase64 = messageType{
		"base64",
	}
)
