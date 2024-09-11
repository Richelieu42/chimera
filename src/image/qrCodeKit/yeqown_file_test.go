package qrCodeKit

import (
	"github.com/yeqown/go-qrcode/writer/standard"
	"testing"
)

func TestEncodeToFile(t *testing.T) {
	err := EncodeToFile("https://www.example.com", "_test0.png")
	if err != nil {
		panic(err)
	}
}

func TestEncodeToFile1(t *testing.T) {
	err := EncodeToFile("https://www.example.com", "_test1.png", standard.WithQRWidth(42))
	if err != nil {
		panic(err)
	}
}
