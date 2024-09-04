package imageKit

import "image/color"

type (
	Info struct {
		ExtName    string      `json:"extName"`
		ColorModel color.Model `json:"colorModel"`

		Width  int `json:"width"`
		Height int `json:"height"`
	}
)
