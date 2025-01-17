package imageKit

import "testing"

func TestToPng(t *testing.T) {
	src := "/Users/richelieu/Desktop/sticker.webp"
	dest := "/Users/richelieu/Desktop/a.png"

	if err := ToPng(src, dest); err != nil {
		panic(err)
	}
}
