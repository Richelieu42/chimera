package imageKit

import (
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/crypto/base64Kit"
	"github.com/richelieu42/go-scales/src/http/httpClientKit"
	"regexp"
)

// GetBase64OfImage (硬盘上的)图片 => base64字符串
/*
参考: golang 将图片生成Base64 https://blog.csdn.net/weixin_40292098/article/details/126029489
*/
func GetBase64OfImage(imagePath string) (string, error) {
	if err := fileKit.AssertExistAndIsFile(imagePath); err != nil {
		return "", err
	}

	imageData, err := fileKit.ReadFile(imagePath)
	if err != nil {
		return "", err
	}
	return EncodeToBase64String(imageData)
}

// GetBase64OfWebImage 网络图片 => base64字符串
/*
参考: golang 将图片生成Base64 https://blog.csdn.net/weixin_40292098/article/details/126029489

@param url e.g."https://img.redocn.com/sheying/20150507/pugongying_4267498.jpg"
*/
func GetBase64OfWebImage(url string) (string, error) {
	// 先获取网络图片的内容
	imageData, err := httpClientKit.Get(url, nil)
	if err != nil {
		return "", err
	}

	return EncodeToBase64String(imageData)
}

func EncodeToBase64String(data []byte) (string, error) {
	mimeType := fileKit.GetMimeType(data)
	switch mimeType {
	case "image/jpeg":
		fallthrough
	case "image/png":
		base64Str := strKit.Format("data:%s;base64,%s", mimeType, base64Kit.EncodeToString(data))
		return base64Str, nil
	default:
		return "", errorKit.Simple("mimeType(%s) isn't supported currently", mimeType)
	}
}

// DecodeFromBase64 图片的base64数据 => 图片的数据（可以直接存储到硬盘上）
/*
@param base64 带不带前缀都无所谓(e.g. "data:image/png;base64,"、"data:image/jpeg;base64,"、"data:image/gif;base64,")
*/
func DecodeFromBase64(base64 []byte) ([]byte, error) {
	// 如果有前缀的话，去掉前缀
	re := regexp.MustCompile(`^data:(.+);base64,`)
	tmp := re.Find(base64)
	length := len(tmp)
	if length > 0 {
		base64 = base64[length:]
	}

	return base64Kit.Decode(base64)
}

// DecodeToImageFile
/*
@param target 要生成的图片的路径
*/
func DecodeToImageFile(base64 []byte, dest string) error {
	data, err := DecodeFromBase64(base64)
	if err != nil {
		return err
	}
	return fileKit.WriteToFile(data, dest)
}
