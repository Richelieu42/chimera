package brotliKit

import "github.com/andybalholm/brotli"

const (
	LevelDefaultCompression = brotli.DefaultCompression

	LevelBestSpeed       = brotli.BestSpeed
	LevelBestCompression = brotli.BestCompression
)

// IsValidLevel
/*
@param level 有效范围: [0, 11]
*/
func IsValidLevel(level int) bool {
	return level >= LevelBestSpeed && level <= LevelBestCompression
}
