package pushKit

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Processor interface {
	ProcessWithGin(ctx *gin.Context)

	Process(w http.ResponseWriter, r *http.Request)
}
