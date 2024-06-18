package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger
/*
@param options 	可以为nil，默认: 输出到os.Stdout、DEBUG级别、人类可读的多行输出
*/
func NewLogger(options ...LoggerOption) (logger *zap.Logger) {
	opts := loadOptions(options...)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = opts.EncodeTime
	encoderConfig.EncodeLevel = opts.EncodeLevel

	var encoder zapcore.Encoder
	if opts.IsOutputTypeConsole() {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, opts.WriteSyncer, opts.Level)
	logger = zap.New(core, zap.WithCaller(opts.Caller))
	return
}
