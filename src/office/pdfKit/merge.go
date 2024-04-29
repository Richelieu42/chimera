package pdfKit

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"io"
)

var (
	Merge func(destFile string, inFiles []string, w io.Writer, conf *model.Configuration, dividerPage bool) error = api.Merge

	MergeAppendFile func(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) error = api.MergeAppendFile

	MergeCreateFile func(inFiles []string, outFile string, dividerPage bool, conf *model.Configuration) error = api.MergeCreateFile
)
