package imageKit

import (
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp" // 兼容webp格式解析
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)
