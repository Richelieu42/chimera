package imageKit

import "image"

var (
	// NewRectangle 返回 image.Rectangle 实例.
	NewRectangle func(x0, y0, x1, y1 int) image.Rectangle = image.Rect
)
