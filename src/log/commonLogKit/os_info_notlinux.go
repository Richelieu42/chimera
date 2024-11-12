//go:build !linux

package commonLogKit

func printUlimitInformation(logger Logger) {
	// do nothing
}

func printOsInformation(logger Logger) {
	// do nothing
}

func printCgroupInfo(logger Logger) {
	// do nothing
}
