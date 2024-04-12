package gzipKit

import (
	"compress/gzip"
)

const (
	LevelNoCompression = gzip.NoCompression

	LevelBestSpeed = gzip.BestSpeed

	LevelBestCompression = gzip.BestCompression
)

// IsValidLevel
/*
@param level 有效范围: [-2, 9]
*/
func IsValidLevel(level int) bool {
	return level >= gzip.HuffmanOnly && level <= gzip.BestCompression
}
