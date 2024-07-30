package reqKit

import (
	"github.com/imroc/req/v3"
	"io"
	"log"
)

var (
	NewLogger func(output io.Writer, prefix string, flag int) req.Logger = req.NewLogger

	NewLoggerFromStandardLogger func(l *log.Logger) req.Logger = req.NewLoggerFromStandardLogger
)
