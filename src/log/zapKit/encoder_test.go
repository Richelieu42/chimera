package zapKit

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestNewEncoder(t *testing.T) {
	enc := NewEncoder()
	core := zapcore.NewCore(enc, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)

	logger.Debug("This is a debug message", zap.String("key", "value"))
	logger.Info("This is an info message")
	logger.Warn("This is a warning message")
	logger.Error("This is an error message0\nThis is an error message1", zap.String("key", "value"), zap.Error(context.Canceled))

	//type args struct {
	//	options []EncoderOption
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want zapcore.Encoder
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := NewEncoder(tt.args.options...); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("NewEncoder() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}
