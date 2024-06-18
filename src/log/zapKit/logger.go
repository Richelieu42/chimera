package zapKit

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger(out zapcore.WriteSyncer, options ...LoggerOption) (logger *zap.Logger, err error) {
	opts := loadOptions(options...)

}
