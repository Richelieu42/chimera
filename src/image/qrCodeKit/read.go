package qrCodeKit

import (
	"github.com/tuotoo/qrcode"
	"os"
)

// Read 解析二维码（仅支持部分图片）.
func Read(path string) (content string, err error) {
	// 打开二维码图片文件
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// 解析二维码
	qrCode, err := qrcode.Decode(file)
	if err != nil {
		return "", err
	}

	return qrCode.Content, nil
}
