package zapKit

import (
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

// 自定义编码器来给日志消息添加前缀
type prefixEncoder struct {
	zapcore.Encoder

	prefix string
}

func (pe *prefixEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 给msg字段加上前缀
	entry.Message = pe.prefix + entry.Message
	return pe.Encoder.EncodeEntry(entry, fields)
}

// NewPrefixEncoder 会给msg字段加上前缀.
/*
@param encoder 不能为nil
*/
func NewPrefixEncoder(encoder zapcore.Encoder, prefix string) zapcore.Encoder {
	if pe, ok := encoder.(*prefixEncoder); ok {
		pe.prefix = prefix
		return pe
	}
	if prefix == "" {
		return encoder
	}

	return &prefixEncoder{
		Encoder: encoder,
		prefix:  prefix,
	}
}
