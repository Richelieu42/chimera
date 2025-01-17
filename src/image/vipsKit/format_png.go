package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"os"
)

// ToPng
/*
@param exportParams 可以为nil
*/
func ToPng(src, dest string, exportParams *vips.PngExportParams) error {
	if err := fileKit.AssertNotExistOrIsFile(dest); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(dest); err != nil {
		return err
	}

	if exportParams == nil {
		exportParams = vips.NewPngExportParams()
		exportParams.Quality = 100
	}

	imageRef, err := Read(src, nil)
	if err != nil {
		return err
	}

	imageData, _, err := imageRef.ExportPng(exportParams)
	if err != nil {
		return err
	}
	return os.WriteFile(dest, imageData, defaultPerm)
}
