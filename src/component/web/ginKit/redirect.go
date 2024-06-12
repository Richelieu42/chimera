package ginKit

import "github.com/gin-gonic/gin"

func Redirect(ctx *gin.Context, code int, location string) {
	ctx.Redirect(code, location)
}
