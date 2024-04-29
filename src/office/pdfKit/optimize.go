package pdfKit

import (
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
	"io"
)

var (
	Optimize func(rs io.ReadSeeker, w io.Writer, conf *model.Configuration) error = api.Optimize

	OptimizeContext func(ctx *model.Context) error = api.OptimizeContext

	OptimizeFile func(inFile, outFile string, conf *model.Configuration) error = api.OptimizeFile
)
