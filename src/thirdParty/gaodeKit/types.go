package gaodeKit

import "github.com/richelieu-yang/chimera/v2/src/core/errorKit"

type (
	BaseResponse struct {
		// Status 值为0或1，1：成功；0：失败
		Status   string `json:"status"`
		InfoCode string `json:"infocode"`
		Info     string `json:"info"`
	}
)

func (baseResp *BaseResponse) IsSuccess() error {
	if baseResp.Status != "1" {
		return errorKit.New("Fail to get weather, status: %s, infocode: %s, info: %s", baseResp.Status, baseResp.InfoCode, baseResp.Info)
	}
	return nil
}
