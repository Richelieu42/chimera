package lz4Kit

import (
	"github.com/pierrec/lz4/v4"
	"io"
)

var (
	NewWriter func(w io.Writer) *lz4.Writer = lz4.NewWriter

	NewReader func(r io.Reader) *lz4.Reader = lz4.NewReader
)
