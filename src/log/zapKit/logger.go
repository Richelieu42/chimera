package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewLogger
/*
@param options 	可以为nil，默认: 	(1) 输出到os.Stdout;
								(2) DEBUG级别;
								(3) 人类可读的多行输出;
								(4) 生产环境(development == false)
*/
func NewLogger(options ...LoggerOption) (logger *zap.Logger) {
	opts := loadOptions(options...)

	/* encoder */
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = opts.EncodeTime
	encoderConfig.EncodeLevel = opts.EncodeLevel
	var encoder zapcore.Encoder
	if opts.IsOutputTypeConsole() {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}
	encoder = NewPrefixEncoder(encoder, opts.MessagePrefix)

	/* core */
	// 第 1 个输出
	core := zapcore.NewCore(encoder, opts.WriteSyncer, opts.LevelEnabler)
	if opts.OtherWriteSyncer != nil && opts.OtherLevelEnabler != nil {
		// 第 2 个输出
		core1 := zapcore.NewCore(encoder, opts.OtherWriteSyncer, opts.OtherLevelEnabler)
		core = zapcore.NewTee(core, core1)
	}
	// initial fields
	if len(opts.InitialFields) > 0 {
		core = core.With(opts.InitialFields)
	}

	/* logger */
	zapOptions := []zap.Option{zap.WithCaller(opts.Caller)}
	if opts.Development {
		zapOptions = append(zapOptions, zap.Development())
	}
	zapOptions = append(zapOptions, zap.AddCallerSkip(opts.CallerSkip))

	logger = zap.New(core, zapOptions...)
	return
}

func NewSugarLogger(options ...LoggerOption) *zap.SugaredLogger {
	logger := NewLogger(options...)
	return logger.Sugar()
}
