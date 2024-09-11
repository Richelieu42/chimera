package main

import (
	_ "image/jpeg"
	_ "image/png"

	"fmt"
	"github.com/tuotoo/qrcode"
	"os"
)

func main() {
	// 打开二维码图片文件
	file, err := os.Open("/Users/richelieu/GolandProjects/chimera/src/image/qrCodeKit/_test-qr2.png")
	if err != nil {
		fmt.Println("打开文件出错:", err)
		return
	}
	defer file.Close()

	// 解析二维码
	qrCode, err := qrcode.Decode(file)
	if err != nil {
		fmt.Println("解析二维码失败:", err)
		return
	}

	// 打印二维码内容
	fmt.Println("二维码内容:", qrCode.Content)
}
