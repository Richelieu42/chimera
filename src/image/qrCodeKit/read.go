package qrCodeKit

import (
	"github.com/tuotoo/qrcode"
	"os"
)

// Read 解析二维码（仅支持部分图片）.
/*
e.g. 不能解析 github.com/skip2/go-qrcode 生成的二维码.
*/
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
