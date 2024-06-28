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

e.g. case: core传nil，options不传
	(1) [Encoder] 人类可读的多行输出
	(2) [Encoder] 时间格式: "2024-06-28T09:15:16.176+0800"
	(3) [Encoder] 日志级别大写且有颜色
	(4) [Encoder] Message字段无前缀
	(5) 仅有1个输出: 输出到控制台(os.Stdout)
	(6) 有 Caller 且 CallerSkip == 0
	(7) Development == false，即生产模式
*/
func NewLogger(core zapcore.Core, options ...LoggerOption) (logger *zap.Logger) {
	if core == nil {
		encoder := NewEncoder()
		// 确保多个goroutine在写入日志时不会发生竞态条件
		ws := zapcore.Lock(os.Stdout)
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
