package service

import (
	"gitee.com/richelieu042/go-scales/src/core/strKit"
	"gitee.com/richelieu042/go-scales/src/jsonKit"
	"gitee.com/richelieu042/go-scales/src/netKit"
)

type (
	Service struct {
		netKit.Address

		// Extra 额外数据
		Extra interface{} `json:"extra,omitempty"`
	}
)

func (service Service) ToJson() (string, error) {
	return jsonKit.MarshalToString(service)
}

func NewService(json string) (*Service, error) {
	if strKit.IsEmpty(json) {
		return nil, nil
	}

	service := &Service{}
	if err := jsonKit.UnmarshalFromString(json, service); err != nil {
		return nil, err
	}
	return service, nil
}
