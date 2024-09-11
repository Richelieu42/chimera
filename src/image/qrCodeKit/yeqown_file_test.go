package qrCodeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/image/imageKit"
	"github.com/yeqown/go-qrcode/writer/standard"
	"testing"
)

func TestEncodeToFile(t *testing.T) {
	err := EncodeToFile("https://www.example.com", "_test0.png")
	if err != nil {
		panic(err)
	}
}

// 带 logo 的二维码（注意，logo的尺寸不能大于二维码的1/5）.
func TestEncodeToFile1(t *testing.T) {
	img, _, err := imageKit.DecodeWithImagePath("_logo.png")
	if err != nil {
		panic(err)
	}

	err = EncodeToFile("https://www.example.com", "_test1.png", standard.WithLogoSizeMultiplier(3), standard.WithLogoImage(img))
	if err != nil {
		panic(err)
	}
}
