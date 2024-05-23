package archiverKit

import (
	"github.com/mholt/archiver/v4"
	"io"
)

var (
	// Identify 识别格式.
	/*
		https://github.com/mholt/archiver?tab=readme-ov-file#identifying-formats
	*/
	Identify func(filename string, stream io.Reader) (archiver.Format, io.Reader, error) = archiver.Identify
)
