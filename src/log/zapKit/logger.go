package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewLogger
/*
@param out 		可以为nil，将输出到os.Stdout
@param options 	可以为nil，默认: DEBUG级别、人类可读的多行输出
*/
func NewLogger(out zapcore.WriteSyncer, options ...LoggerOption) (logger *zap.Logger) {
	if out == nil {
		out = zapcore.AddSync(os.Stdout)
	}
	opts := loadOptions(options...)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	if opts.IsOutputTypeConsole() && opts.ColorWhenConsoleEncoder {
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	}

	var encoder zapcore.Encoder
	if opts.IsOutputTypeConsole() {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	core := zapcore.NewCore(encoder, out, opts.Level)
	logger = zap.New(core, zap.WithCaller(opts.Caller))
	return
}
