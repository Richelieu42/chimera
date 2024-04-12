package gzipKit

// CompressComplexly
/*
PS: 不涉及 compressThreshold 的话，建议直接使用 Compress.
*/
func CompressComplexly(data []byte, options ...GzipOption) ([]byte, error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}
