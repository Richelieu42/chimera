package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewEncoder
/*
默认（不传参）:
	(1) [Encoder] 人类可读的多行输出
	(2) [Encoder] 时间格式: "2024-06-28T09:15:16.176+0800"
	(3) [Encoder] 日志级别大写且有颜色
	(4) [Encoder] Message字段无前缀
*/
func NewEncoder(options ...EncoderOption) zapcore.Encoder {
	opts := loadEncoderOptions(options...)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = opts.EncodeTime
	encoderConfig.EncodeLevel = opts.EncodeLevel

	var encoder zapcore.Encoder
	if opts.IsOutputFormatConsole() {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	encoder = NewPrefixEncoder(encoder, opts.MessagePrefix)

	return encoder
}
