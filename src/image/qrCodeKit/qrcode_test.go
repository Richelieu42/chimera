package qrCodeKit

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"testing"
)

func TestWriteFile(t *testing.T) {
	content := "https://example.org"

	err := WriteFile(content, qrcode.Medium, 256, "_test-qr0.png")
	if err != nil {
		panic(err)
	}
}

func TestWriteColorFile(t *testing.T) {
	content := "https://example.org"

	err := WriteColorFile(content, qrcode.Medium, 256, color.Transparent, color.Black, "_test-qr1.png")
	if err != nil {
		panic(err)
	}
}

func TestWriteFileWithBackgroundImage(t *testing.T) {
	content := "https://example.org"
	bgPath := "/Users/richelieu/Desktop/iShot_2024-09-04_13.51.58.PNG"

	err := WriteFileWithBackgroundImage(content, qrcode.Medium, 1000, bgPath, color.Black, "_test-qr2.png")
	if err != nil {
		panic(err)
	}
}
