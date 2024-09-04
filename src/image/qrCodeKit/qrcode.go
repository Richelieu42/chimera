package qrCodeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
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
/*
PS: 背景色默认为白色（非透明），前景色默认为黑色.

@param size 生成图片的尺寸
			e.g.256 => 256*256
@param outputImagePath 	输出图片的路径
						(1) 如果存在且是个文件的话，会覆盖
						(2) 建议是 .png 格式的
						(3) 生成图片的背景色是白色而非透明，即使保存为 .png 格式
*/
func WriteFile(content string, level qrcode.RecoveryLevel, size int, outputImagePath string) error {
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}

	return qrcode.WriteFile(content, level, size, outputImagePath)
}

// WriteColorFile
/*
@param background 背景色（推荐使用透明色 color.Transparent，然后保存为.png格式的图片）
@param foreground 前景色（一般为 color.Black）
*/
func WriteColorFile(content string, level qrcode.RecoveryLevel, size int, background, foreground color.Color, outputImagePath string) error {
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}

	return qrcode.WriteColorFile(content, level, size, background, foreground, outputImagePath)
}

func WriteFileWithBackgroundImage(content string, level qrcode.RecoveryLevel, size int, backgroundImagePath string, foreground color.Color, outputImagePath string) error {
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}

	return nil
}
