package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mholt/archiver/v4"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	_ "github.com/richelieu-yang/chimera/v3/src/log/logrusInitKit"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	engine := gin.Default()

	engine.Any("/test", func(ctx *gin.Context) {
		logrus.Info("---")
		err := pack(ctx.Request.Context())
		logrus.Info("===")
		if err != nil {
			logrus.WithError(err).Error("fail to pack")
			ctx.String(200, err.Error())
			return
		}
		ctx.String(200, "ok")
	})

	if err := engine.Run(":80"); err != nil {
		logrus.Fatal(err)
	}
}

func pack(ctx context.Context) error {
	// map files on disk to their paths in the archive
	files, err := archiver.FilesFromDisk(nil, map[string]string{
		//"/path/on/disk/file1.txt":       "file1.txt",
		//"/path/on/disk/file2.txt":       "subfolder/file2.txt",
		//"/path/on/disk/file3.txt":       "",              // put in root of archive as file3.txt
		//"/path/on/disk/file4.txt":       "subfolder/",    // put in subfolder as file4.txt
		"/Users/richelieu/Documents/ino/notes": "Custom Folder", // contents added recursively
	})
	if err != nil {
		return err
	}

	// create the output file we'll write to
	out, err := os.Create(idKit.NewUUID() + ".tar.gz")
	if err != nil {
		return err
	}
	defer out.Close()

	// we can use the CompressedArchive type to gzip a tarball
	// (compression is not required; you could use Tar directly)
	format := archiver.CompressedArchive{
		Compression: archiver.Gz{},
		Archival:    archiver.Tar{},
	}

	// create the archive
	err = format.Archive(ctx, out, files)
	if err != nil {
		return err
	}
	return nil
}
