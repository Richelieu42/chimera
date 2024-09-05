package imageKit

import "image"

var (
	NewPointer func(X, Y int) image.Point = image.Pt
)
