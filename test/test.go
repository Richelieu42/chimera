package main

import (
	"github.com/yeqown/go-qrcode/writer/standard"
	"image"
	"log"
	"os"

	"github.com/yeqown/go-qrcode/v2"
)

func main() {
	// 要编码的内容
	content := "https://example.com"

	// 生成二维码(*qrcode.QRCode实例)
	qr, err := qrcode.New(content)
	if err != nil {
		panic(err)
	}

	// 打开 Logo 图片文件
	logoFile, err := os.Open("logo.png")
	if err != nil {
		log.Fatalf("无法打开 Logo 文件: %v", err)
	}
	defer logoFile.Close()

	// 读取 Logo 图片
	logo, _, err := image.Decode(logoFile)
	if err != nil {
		log.Fatalf("无法解码 Logo 文件: %v", err)
	}

	// 创建二维码图片并将 Logo 添加到二维码中央
	qrWriter, err := standard.New("./output_with_logo.png",
		standard.WithLogoImage(logo), // 添加 Logo
		standard.WithLogoSizeMultiplier(2),
	)
	if err != nil {
		panic(err)
	}

	// 保存二维码到文件
	if err := qr.Save(qrWriter); err != nil {
		log.Fatalf("保存二维码失败: %v", err)
	}

	log.Println("二维码生成成功，并保存为 output_with_logo.png")
}
