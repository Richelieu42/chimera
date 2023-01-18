// Package ginKit
/*
TODO: 响应文件或文件流，目前一律用no-cache，在此种情况下会有问题：一个网址对应复数的图片（可以参考Web.docx）.
*/
package ginKit

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/richelieu42/go-scales/src/consts/httpStatusCode"
	"github.com/richelieu42/go-scales/src/core/file/fileKit"
	"github.com/richelieu42/go-scales/src/core/intKit"
	"github.com/richelieu42/go-scales/src/core/strKit"
	"github.com/richelieu42/go-scales/src/jsonKit"
	"net/http"
	"net/url"
	"unicode"
)

// RespondText 响应文本.
/*
@param args 用于替换 传参format 中的格式占位符（%s、%d...）
*/
func RespondText(ctx *gin.Context, statusCode int, format string, args ...interface{}) {
	SetCacheControlNoCache(ctx)

	ctx.String(statusCode, format, args...)
}

// RespondJson 响应json字符串.
/*
@param h gin.H is a shortcut for map[string]interface{}
*/
func RespondJson(ctx *gin.Context, statusCode int, h gin.H) {
	SetCacheControlNoCache(ctx)

	ctx.JSON(statusCode, h)
}

// RespondFile
/*
PS:
(1) 支持的浏览器: ie、Edge、Safari、Chrome、Firefox；
(2) 按照实际的业务，自行判断是否和 ginKit.SetCacheControlNoCache() 或 ginKit.SetCacheControlNoStore() 组合使用，或者不设置 "Cache-Control".

@param name 文件名（如果为""，将自动从传参path中获取）
*/
func RespondFile(ctx *gin.Context, httpStatusCode int, path, name string) {
	ctx.Status(httpStatusCode)

	if strKit.IsEmpty(name) {
		ctx.File(path)
	} else {
		// 内部会对文件名进行处理，因此此处不用再额外处理 name
		ctx.FileAttachment(path, name)
	}
}

// RespondFileContent 响应文件流([]byte)给客户端.
/*
PS:
(1) 支持的浏览器: ie、Edge、Safari、Chrome、Firefox；
(2) 按照实际的业务，自行判断是否和 ginKit.SetCacheControlNoCache() 或 ginKit.SetCacheControlNoStore() 组合使用，或者不设置 "Cache-Control".

@param name 		可以为""
@param contentType 	可以为""，将会使用默认值
*/
func RespondFileContent(ctx *gin.Context, httpStatusCode int, name, contentType string, data []byte) {
	ctx.Status(httpStatusCode)

	// https://www.runoob.com/http/http-content-type.html
	// application/octet-stream: 二进制流数据（如常见的文件下载）
	contentType = strKit.EmptyToDefault(contentType, "application/octet-stream; charset=utf-8", true)
	ctx.Header("Content-Type", contentType)

	ctx.Header("Content-Length", intKit.IntToString(len(data)))

	if strKit.IsNotEmpty(name) {
		/* 参考: ctx.FileAttachment()，以防中文文件名乱码 */
		isASCII := func(s string) bool {
			for i := 0; i < len(s); i++ {
				if s[i] > unicode.MaxASCII {
					return false
				}
			}
			return true
		}
		if isASCII(name) {
			ctx.Header("Content-Disposition", `attachment; filename="`+name+`"`)
		} else {
			ctx.Header("Content-Disposition", `attachment; filename*=UTF-8''`+url.QueryEscape(name))
		}
	}

	_, _ = ctx.Writer.Write(data)
}

// RespondFileFromFS
/*
PS:
(1) 支持的浏览器: ie、Edge、Safari、Chrome、Firefox；
(2) 按照实际的业务，自行判断是否和 ginKit.SetCacheControlNoCache() 或 ginKit.SetCacheControlNoStore() 组合使用，或者不设置 "Cache-Control".
*/
func RespondFileFromFS(ctx *gin.Context, httpStatusCode int, filePath string, fs http.FileSystem) {
	ctx.Status(httpStatusCode)

	ctx.FileFromFS(filePath, fs)
}

func RespondIcon(ctx *gin.Context, httpStatusCode int, iconData []byte) {
	ctx.Data(httpStatusCode, "image/x-icon; charset=UTF-8", iconData)
}

// RespondHtml 响应html（[]byte形式）给前端.
/*
PS:
(1) 只能是单纯的html文件，就算有js和css也只能内嵌，不能从外部导入；
(2) 无法进行渲染（ctx.HTML()可以进行渲染）；
(3) 可以搭配 go-bindata 一起使用.
*/
func RespondHtml(ctx *gin.Context, httpStatusCode int, htmlData []byte) {
	ctx.Data(httpStatusCode, "text/html; charset=UTF-8", htmlData)
}

// RespondPdfContentToPrint 会触发浏览器端的打印.
/*
PS:
(1) 不要传文件名
(2) contentType为"application/pdf"
*/
func RespondPdfContentToPrint(ctx *gin.Context, httpStatusCode int, pdfContent []byte) {
	RespondFileContent(ctx, httpStatusCode, "", "application/pdf; charset=UTF-8", pdfContent)
}

func RespondError(ctx *gin.Context, httpStatusCode int, err error) {
	var message string
	if err != nil {
		message = err.Error()
	} else {
		message = "err == nil"
	}

	httpStatusCode = handleHttpStatusCode(httpStatusCode)
	json, _ := jsonKit.SealFully(nil, "error", message, nil)
	RespondText(ctx, httpStatusCode, json)
}

func RespondPanic(ctx *gin.Context, err any) {
	var message string
	if err != nil {
		message = fmt.Sprintf("%v", err)
	} else {
		message = "err == nil"
	}

	json, _ := jsonKit.SealFully(nil, "panic", message, nil)
	RespondText(ctx, httpStatusCode.Panic, json)
}

func RespondPackageOrError(ctx *gin.Context, pack *ResponsePackage, err error) {
	if err != nil {
		RespondError(ctx, http.StatusOK, err)
		return
	}
	RespondPackage(ctx, pack)
}

// RespondPackage
/*
PS:
优先顺序（从高到低）：文本（包括json）、 文件（路径）、文件（内容）、错误(error)

@param pack 可以为nil，此时：状态码为200，响应内容为空（不存在请求转发的情况）.
*/
func RespondPackage(ctx *gin.Context, pack *ResponsePackage) {
	if pack == nil {
		// 可能的场景：请求转发成功了，不用处理了；响应内容为空
		return
	}

	// 处理"零值"的情况
	statusCode := handleHttpStatusCode(pack.StatusCode)

	// (1) 文本（包括json）
	if strKit.IsNotEmpty(pack.Text) {
		RespondText(ctx, statusCode, pack.Text)
		return
	}
	// (2) 文件（路径）
	if strKit.IsNotEmpty(pack.FilePath) {
		if strKit.IsEmpty(pack.FileName) {
			pack.FileName = fileKit.GetName(pack.FilePath)
		}
		RespondFile(ctx, statusCode, pack.FilePath, pack.FileName)
		return
	}
	// (3) 文件（内容）
	if pack.FileContent != nil {
		RespondFileContent(ctx, statusCode, pack.FileName, pack.ContentType, pack.FileContent)
		return
	}
	// (4) 错误(error)
	if pack.Error != nil {
		RespondError(ctx, statusCode, pack.Error)
		return
	}

	ctx.Status(statusCode)
}

func handleHttpStatusCode(code int) int {
	if code <= 0 {
		code = http.StatusOK
	}
	return code
}
