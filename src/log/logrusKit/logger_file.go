package logrusKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/ioKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/sirupsen/logrus"
	"os"
)

// NewFileLogger 输出到 文件(not rotatable).
/*
@param filePath (1) 如果文件不存在，会自动创建；
				(2) 如果文件存在，会自动追加内容.
*/
func NewFileLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	file, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	options = append(options, WithOutput(file))

	return NewLogger(options...), nil
}

// NewFileAndStdoutLogger 同时输出到 文件(not rotatable) 和 os.Stdout.
/*
@param filePath (1) 如果文件不存在，会自动创建；
				(2) 如果文件存在，会自动追加内容.
*/
func NewFileAndStdoutLogger(filePath string, options ...LoggerOption) (*logrus.Logger, error) {
	f, err := fileKit.CreateInAppendMode(filePath)
	if err != nil {
		return nil, err
	}
	output := ioKit.MultiWriter(f, os.Stdout)
	options = append(options, WithOutput(output))

	return NewLogger(options...), nil
}
