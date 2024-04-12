package gzipKit

func Compress(data []byte, options ...GzipOption) ([]byte, error) {
	opts := loadOptions(options...)

	return opts.Compress(data)
}
