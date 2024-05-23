package archiverKit

import (
	"compress/gzip"
	"github.com/mholt/archiver/v4"
)

var (
	zipCompressedArchive = &archiver.CompressedArchive{
		Archival: archiver.Zip{},
	}

	tarGzCompressedArchive = &archiver.CompressedArchive{
		Compression: archiver.Gz{
			CompressionLevel: gzip.BestSpeed,
		},
		Archival: archiver.Tar{},
	}
)
