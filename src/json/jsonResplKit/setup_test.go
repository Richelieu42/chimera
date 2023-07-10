package jsonResplKit

import (
	"fmt"
	"github.com/richelieu-yang/chimera/v2/src/consts"
	"github.com/richelieu-yang/chimera/v2/src/core/pathKit"
	"testing"
)

func TestMustSetUp(t *testing.T) {
	if _, err := pathKit.ReviseWorkingDirInTestMode(consts.ProjectName); err != nil {
		panic(err)
	}

	type Response struct {
		ErrorCode    string      `json:"errorCode"`
		ErrorMessage string      `json:"errorMessage"`
		Result       interface{} `json:"result,omitempty"`
	}

	provider := func(code, msg string, data interface{}) interface{} {
		return &Response{
			ErrorCode:    code,
			ErrorMessage: msg,
			Result:       data,
		}
	}
	MustSetUp(provider, []string{"chimera-lib/msg.properties"})

	fmt.Println(Seal("0", nil))
	fmt.Println(Seal("42", 666, "tester"))
}
