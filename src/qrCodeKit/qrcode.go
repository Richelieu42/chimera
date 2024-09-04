package qrCodeKit

import (
	"github.com/skip2/go-qrcode"
	"image/color"
)

// Encode 生成二维码([]byte类型).
/*
@param level 一般使用 qrcode.Medium
*/
func Encode(content string, level qrcode.RecoveryLevel, size int) ([]byte, error) {
	return qrcode.Encode(content, level, size)
}

// WriteFile 生成二维码文件.
func WriteFile(content string, level qrcode.RecoveryLevel, size int, filename string) error {
	return qrcode.WriteFile(content, level, size, filename)
}

func WriteColorFile(content string, level qrcode.RecoveryLevel, size int, background, foreground color.Color, filename string) error {
	return qrcode.WriteColorFile(content, level, size, background, foreground, filename)
}
