//go:build !linux

package commonLogKit

func printLimitsForCurrentPid(logger Logger) {
	// do nothing
}

func printOsInfo(logger Logger) {
	// do nothing
}

func printUlimitInfo(logger Logger) {
	// do nothing
}

func printCgroupInfo(logger Logger) {
	// do nothing
}
