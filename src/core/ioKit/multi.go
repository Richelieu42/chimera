package ioKit

import (
	"io"
)

var (
	MultiReader func(readers ...io.Reader) io.Reader = io.MultiReader

	MultiWriter func(writers ...io.Writer) io.Writer = io.MultiWriter
)
