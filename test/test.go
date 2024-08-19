package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
	"os"
)

func main() {
	// 定义输入 PDF 文件和输出目录
	inputPDF := "/Users/richelieu/Desktop/34页(1).pdf"
	outputDir := "_images"

	// 创建输出目录
	err := os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		log.Fatalf("无法创建输出目录: %v", err)
	}

	//// 打开 PDF 文件
	//pdfFile, err := os.Open(inputPDF)
	//if err != nil {
	//	log.Fatalf("无法打开 PDF 文件: %v", err)
	//}
	//defer pdfFile.Close()

	// 提取图像
	err = api.ExtractImagesFile(inputPDF, outputDir, []string{"1", "2"}, nil)
	if err != nil {
		log.Fatalf("提取图像失败: %v", err)
	}

	fmt.Printf("图像提取完成，保存在目录: %s\n", outputDir)
}
