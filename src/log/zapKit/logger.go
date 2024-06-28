package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// NewLogger
/*
@param core		可以为nil
@param options	可以不传
*/
func NewLogger(core zapcore.Core, options ...LoggerOption) (logger *zap.Logger) {
	if core == nil {
		encoder := NewEncoder()
		ws := zapcore.AddSync(os.Stdout)
		core = NewCore(encoder, ws, zapcore.DebugLevel)
	}

	opts := loadOptions(options...)

	var zapOptions []zap.Option
	// Development
	if opts.Development {
		zapOptions = append(zapOptions, zap.Development())
	}
	// Caller
	zapOptions = append(zapOptions, zap.WithCaller(opts.Caller))
	// CallerSkip
	zapOptions = append(zapOptions, zap.AddCallerSkip(opts.CallerSkip))

	logger = zap.New(core, zapOptions...)
	return
}

func NewSugarLogger(core zapcore.Core, options ...LoggerOption) *zap.SugaredLogger {
	logger := NewLogger(core, options...)
	return logger.Sugar()
}
