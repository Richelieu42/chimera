package main

import (
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"log"
)

func mergePDFs(inputFiles []string, outputFile string) error {
	return api.MergeCreateFile(inputFiles, outputFile, false, nil)
}

func main() {
	files := []string{"/Users/richelieu/Desktop/a.pdf", "/Users/richelieu/Desktop/b.pdf"}
	output := "merged1.pdf"

	err := mergePDFs(files, output)
	if err != nil {
		log.Fatalf("Error merging PDFs: %v", err)
	}

	fmt.Println("PDFs merged successfully!")
}
