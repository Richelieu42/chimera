package ioKit

import "io"

// Copy
var Copy func(dst io.Writer, src io.Reader) (written int64, err error) = io.Copy

// CopyN
var CopyN func(dst io.Writer, src io.Reader, n int64) (written int64, err error) = io.CopyN

// CopyBuffer
var CopyBuffer func(dst io.Writer, src io.Reader, buf []byte) (written int64, err error) = io.CopyBuffer
