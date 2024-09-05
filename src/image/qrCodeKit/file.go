package qrCodeKit

import (
	"github.com/richelieu-yang/chimera/v3/src/core/errorKit"
	"github.com/richelieu-yang/chimera/v3/src/core/mathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/pathKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"github.com/richelieu-yang/chimera/v3/src/file/fileKit"
	"github.com/richelieu-yang/chimera/v3/src/idKit"
	"github.com/richelieu-yang/chimera/v3/src/image/imageKit"
	"github.com/skip2/go-qrcode"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

// WriteFile 生成二维码文件.
/*
PS: 背景色默认为白色（非透明），前景色默认为黑色.

@param size 生成图片的尺寸
			e.g.256 => 256*256
@param outputImagePath 	输出图片的路径
						(1) 如果存在且是个文件的话，会覆盖
						(2) 建议是 .png 格式的
						(3) 生成图片的背景色是白色而非透明，即使保存为 .png 格式
*/
func WriteFile(content string, level qrcode.RecoveryLevel, size int, outputImagePath string) error {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return err
	}
	/* outputImagePath */
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}

	return qrcode.WriteFile(content, level, size, outputImagePath)
}

// WriteFileWithColor
/*
@param background 背景色（推荐使用透明色 color.Transparent，然后保存为.png格式的图片）
@param foreground 前景色（一般为 color.Black）
@param outputImagePath 输出的图片路径，仅支持3种格式: .jpg、.jpeg、.png（推荐）
*/
func WriteFileWithColor(content string, level qrcode.RecoveryLevel, size int, background, foreground color.Color, outputImagePath string) error {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return err
	}
	/* outputImagePath */
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}

	return qrcode.WriteColorFile(content, level, size, background, foreground, outputImagePath)
}

// WriteFileWithBackgroundImage
/*
@param size 二维码的尺寸，如果<=0，则自适应（取背景图片宽高的最小值）
*/
func WriteFileWithBackgroundImage(content string, level qrcode.RecoveryLevel, size int, backgroundImagePath string, foreground color.Color, outputImagePath string) error {
	/* content */
	if err := strKit.AssertNotEmpty(content, "content"); err != nil {
		return err
	}
	/* outputImagePath */
	if err := fileKit.AssertNotExistOrIsFile(outputImagePath); err != nil {
		return err
	}
	if err := fileKit.MkParentDirs(outputImagePath); err != nil {
		return err
	}
	outputExt := fileKit.GetExt(outputImagePath)
	var jpgFlag bool
	switch outputExt {
	case ".png":
		jpgFlag = false
	case ".jpg", ".jpeg":
		jpgFlag = true
	default:
		return errorKit.Newf("invalid outputExt(%s)", outputExt)
	}

	/* backgroundImagePath */
	if err := fileKit.AssertExistAndIsFile(backgroundImagePath); err != nil {
		return err
	}
	bgFile, err := os.Open(backgroundImagePath)
	if err != nil {
		return err
	}
	bgImg, _, err := imageKit.Decode(bgFile)
	if err != nil {
		return err
	}
	bgBounds := bgImg.Bounds()
	bgWidth := bgBounds.Dx()
	bgHeight := bgBounds.Dy()

	/* size */
	if size <= 0 {
		size = mathKit.Min(bgWidth, bgHeight)
	}

	/* 重新绘制 */
	width := mathKit.Max(size, bgWidth)
	height := mathKit.Max(size, bgHeight)
	bounds := imageKit.NewRectangle(0, 0, width, height)
	img := image.NewRGBA(bounds)

	/* (1) 生成二维码文件 */
	var qrImg image.Image
	{
		dirPath := pathKit.ParentDir(outputImagePath)
		qrPath := pathKit.Join(dirPath, "_"+idKit.NewXid()+".png")
		if err := WriteFileWithColor(content, level, size, color.Transparent, foreground, qrPath); err != nil {
			return err
		}
		defer fileKit.Delete(qrPath)

		f, err := os.Open(qrPath)
		if err != nil {
			return err
		}
		defer f.Close()
		qrImg, _, err = imageKit.Decode(f)
		if err != nil {
			return err
		}
	}

	/* (1.5) 输出为.jpg或.jpeg格式的情况下，需要先画一层白色底色（否则如果背景图片中有透明色的话，那部分会变成黑色） */
	if jpgFlag {
		draw.Draw(img, bounds, &image.Uniform{C: color.White}, image.Point{}, draw.Over)
	}

	/* (2) 先绘制背景图片 */
	{
		pt := image.Pt((width-bgWidth)/2, (height-bgHeight)/2)
		tmpBounds := bgBounds.Add(pt)

		draw.Draw(img, tmpBounds, bgImg, image.Point{}, draw.Over)
	}

	/* (3) 再绘制二维码 */
	{
		pt := image.Pt((width-size)/2, (height-size)/2)
		tmpBounds := qrImg.Bounds().Add(pt)

		draw.Draw(img, tmpBounds, qrImg, image.Point{}, draw.Over)
	}

	/* (4) 将合成后的图片保存为新文件 */
	outFile, err := os.Create(outputImagePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	if jpgFlag {
		return jpeg.Encode(outFile, img, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		})
	}
	return png.Encode(outFile, img)
}
