package vipsKit

import (
	"github.com/davidbyttow/govips/v2/vips"
	"github.com/sirupsen/logrus"
)

// SetUp
/*
PS: vips 8.14.5支持的格式: webp、heif、pdf、png、tiff、svg、jp2k、gif、jpeg、magick.

@param config 可以为nil（使用默认配置: concurrency=1 cache_max_files=0 cache_max_mem=52428800 cache_max=100）
*/
func SetUp(config *vips.Config) {
	// 会输出一些信息到控制台
	vips.Startup(config)

	logrus.RegisterExitHandler(func() {
		vips.Shutdown()
	})
}
