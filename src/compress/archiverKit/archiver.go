package archiverKit

import (
	"context"
	"github.com/mholt/archiver/v4"
	"github.com/richelieu-yang/chimera/v3/src/core/interfaceKit"
	"github.com/richelieu-yang/chimera/v3/src/core/mapKit"
	"io"
)

// archive 压缩.
/*
@param ctx		可用于中断压缩
				e.g. 请求被取消
					gin中的 ctx.Request.Context() 作为传参
@param format	e.g.
					&archiver.CompressedArchive{
						Compression: archiver.Gz{},
						Archival:    archiver.Tar{},
					}
@param output	如果返回值非nil（可能原因: 传参ctx被取消），则表示压缩失败，应当将生成的压缩文件删掉.
@param mapper	(1) 键:	要压缩文件（或目录）的路径
				(2) 值:	放到压缩包里面的哪里
				(3) e.g.
					map[string]string{
						"/path/on/disk/file1.txt": "file1.txt",
						"/path/on/disk/file2.txt": "subfolder/file2.txt",
						"/path/on/disk/file3.txt": "",              // put in root of archive as file3.txt
						"/path/on/disk/file4.txt": "subfolder/",    // put in subfolder as file4.txt
						"/path/on/disk/folder":    "Custom Folder", // contents added recursively
					}
@param options	可以为nil

*/
func archive(ctx context.Context, format archiver.CompressedArchive, output io.Writer, mapper map[string]string, options *archiver.FromDiskOptions) error {
	if ctx == nil {
		ctx = context.TODO()
	}
	if err := interfaceKit.AssertNotNil(format, "format"); err != nil {
		return err
	}
	if err := interfaceKit.AssertNotNil(output, "output"); err != nil {
		return err
	}
	if err := mapKit.AssertNotEmpty(mapper, "mapper"); err != nil {
		return err
	}

	files, err := archiver.FilesFromDisk(options, mapper)
	if err != nil {
		return err
	}
	return format.Archive(ctx, output, files)
}
