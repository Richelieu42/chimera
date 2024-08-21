package main

import (
	"github.com/richelieu-yang/chimera/v3/src/office/pdfKit"
)

func main() {
	inputPDF := "/Users/richelieu/Desktop/34页(1).pdf"
	outputDir := "_tmp"
	if err := pdfKit.SplitByPageNrFile(inputPDF, outputDir, []int{2}, nil); err != nil {
		panic(err)
	}
}
