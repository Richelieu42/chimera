package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/richelieu-yang/chimera/v2/src/component/web/ginKit"
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit"
	"github.com/richelieu-yang/chimera/v2/src/json/jsonRespKit"
	"net/http"
)

// PushToAll
// @Summary 推送消息给所有连接.
// @Description 推送消息给所有连接（exceptBsids对应的链接例外）.
// @Router /push_to_all [post]
// @Accept application/x-www-form-urlencoded
// @Param text 			body	string		true	"推送消息的内容."
// @Param exceptBsids	body	[]string 	false	"例外连接的bsid."
// @Produce json
func PushToAll(ctx *gin.Context) (*ginKit.ResponsePackage, error) {
	type Params struct {
		Text        string   `form:"text" binding:"required"`
		ExceptBsids []string `form:"exceptBsids,optional"`
	}
	params := &Params{}
	if err := ctx.ShouldBind(params); err != nil {
		return &ginKit.ResponsePackage{
			StatusCode: http.StatusBadRequest,
			Text:       err.Error(),
		}, nil
	}

	err := pushKit.PushToAll([]byte(params.Text), params.ExceptBsids)
	if err != nil {
		return nil, err
	}
	return &ginKit.ResponsePackage{
		Object: jsonRespKit.PackFully("0", "ok", nil),
	}, nil
}
