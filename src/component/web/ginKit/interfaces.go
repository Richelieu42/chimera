package ginKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type (
	IGroup interface {
		Handle(httpMethod, relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

		Any(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes

		StaticFile(relativePath, filepath string) gin.IRoutes

		StaticFS(relativePath string, fs http.FileSystem) gin.IRoutes

		Static(relativePath, root string) gin.IRoutes
	}
)
