package logrusKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/sirupsen/logrus"
	"os"
)

// NewLogger
/*
PS:
(1) 默认输出到 控制台(os.Stderr);
(2) 如果希望 输出到文件 且 rotatable，可以使用 WithOutput()，详见下例.

@param options 可以什么都不配置（此时输出到控制台）
*/
func NewLogger(options ...LoggerOption) *logrus.Logger {
	opts := loadOptions(options...)

	logger := logrus.New()
	logger.SetFormatter(opts.formatter)
	logger.SetReportCaller(opts.reportCaller)
	logger.SetLevel(opts.level)
	if opts.output != nil {
		logger.SetOutput(opts.output)
	}

	// msgPrefix
	if strKit.IsNotEmpty(opts.msgPrefix) {
		hook := &defaultPrefixHook{prefix: opts.msgPrefix}
		logger.AddHook(hook)
	}

	if opts.disableQuote {
		DisableQuote(logger)
	}

	return logger
}

// NewFileLogger 输出到 文件(not rotatable).
func NewFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	file, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	options = append(options, WithOutput(file))

	return NewLogger(options...), nil
}

// NewFileAndStdoutLogger 同时输出到 文件(not rotatable) 和 os.Stdout.
func NewFileAndStdoutLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	f, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	output := ioKit.MultiWriter(f, os.Stdout)
	options = append(options, WithOutput(output))

	return NewLogger(options...), nil
}

// DisposeLogger 释放资源（主要针对文件日志）
func DisposeLogger(logger *logrus.Logger) error {
	if logger == nil {
		return nil
	}

	return ioKit.TryToClose(logger.Out)
}
