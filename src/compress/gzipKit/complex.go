package gzipKit

func CompressComplexly(data []byte, options ...GzipOption) ([]byte, error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}
