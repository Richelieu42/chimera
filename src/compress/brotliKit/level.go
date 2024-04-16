package brotliKit

import "github.com/andybalholm/brotli"

const (
	LevelDefaultCompression = brotli.DefaultCompression

	LevelBestSpeed       = brotli.BestSpeed
	LevelBestCompression = brotli.BestCompression
)

// IsValidLevel
/*
@param level 有效范围: [-2, 9]
*/
func IsValidLevel(level int) bool {
	return level >= LevelBestSpeed && level <= LevelBestCompression
}
