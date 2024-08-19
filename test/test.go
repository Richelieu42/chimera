package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/richelieu-yang/chimera/v3/src/netKit"
	"log"
)

func main() {
	port := 8000
	engine := gin.Default()
	engine.Any("/test", func(ctx *gin.Context) {
		ctx.String(200, fmt.Sprintf("[%d] Hello world!", port))
	})
	if err := engine.Run(netKit.JoinToHost("", port)); err != nil {
		panic(err)
	}
}

func mergePDFs(inputFiles []string, outputFile string) error {
	api.ExtractImages()

	err := api.MergeCreateFile(inputFiles, outputFile, nil)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	files := []string{"file1.pdf", "file2.pdf", "file3.pdf"}
	output := "merged.pdf"

	err := mergePDFs(files, output)
	if err != nil {
		log.Fatalf("Error merging PDFs: %v", err)
	}

	fmt.Println("PDFs merged successfully!")
}
