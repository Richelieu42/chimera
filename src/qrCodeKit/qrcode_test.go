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
