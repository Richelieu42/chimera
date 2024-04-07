package ginKit

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v3/src/core/boolKit"
	"github.com/richelieu-yang/chimera/v3/src/core/floatKit"
	"github.com/richelieu-yang/chimera/v3/src/core/intKit"
	"github.com/richelieu-yang/chimera/v3/src/core/strKit"
	"io"
)

// ObtainGetParam 获取 GET 请求的参数（从url获取）.
/*
PS:
(1) 原生方法的用法更加丰富!
(2) 不需要额外手动解码.
*/
func ObtainGetParam(ctx *gin.Context, key string) string {
	return ctx.Query(key)
}

// ObtainPostParam 获取 POST 请求的参数.
/*
PS:
(1) 原生方法的用法更加丰富!
(2) 不需要额外手动解码;
(3) 请求的Content-Type，	(a) 支持: multipart/form-data、x-www-form-urlencoded（即application/x-www-form-urlencoded）
						(b) 不支持: application/json
*/
func ObtainPostParam(ctx *gin.Context, key string) string {
	return ctx.PostForm(key)
}

// ObtainParam 获取请求参数.
/*
PS:
(1) 优先级（从左到右）: GET、POST;
(2) 不需要额外手动解码.
*/
func ObtainParam(ctx *gin.Context, key string) string {
	// (1) GET
	value := ctx.Query(key)
	if strKit.IsNotEmpty(value) {
		return value
	}

	// (2) POST
	return ctx.PostForm(key)
}

func ObtainBoolParam(ctx *gin.Context, key string) (bool, error) {
	value := ObtainParam(ctx, key)
	return boolKit.ToBoolE(value)
}

func ObtainIntParam(ctx *gin.Context, key string) (int, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToIntE(value)
}

func ObtainInt32Param(ctx *gin.Context, key string) (int32, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToInt32E(value)
}

func ObtainInt64Param(ctx *gin.Context, key string) (int64, error) {
	value := ObtainParam(ctx, key)
	return intKit.ToInt64E(value)
}

func ObtainFloat32Param(ctx *gin.Context, key string) (float32, error) {
	value := ObtainParam(ctx, key)
	return floatKit.ToFloat32E(value)
}

func ObtainFloat64Param(ctx *gin.Context, key string) (float64, error) {
	value := ObtainParam(ctx, key)
	return floatKit.ToFloat64E(value)
}

// ObtainFormFileContent form请求，根据 传参key 获取文件的字节流.（单文件上传）
/*
@return 文件内容 + 文件名 + 错误
*/
func ObtainFormFileContent(ctx *gin.Context, key string) ([]byte, string, error) {
	fileHeader, err := ctx.FormFile(key)
	if err != nil {
		return nil, "", err
	}

	file, err := fileHeader.Open()
	if err != nil {
		return nil, "", err
	}
	defer file.Close()
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, "", err
	}
	return content, fileHeader.Filename, nil
}
