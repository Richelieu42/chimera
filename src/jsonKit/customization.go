package jsonKit

type (
	MsgProcessor func(code string, msg string, data interface{}) string

	JsonResponseProcessor func(resp *JsonResponse) any
)

// 如果不为nil的情况，将根据 code 和 msg 返回一个新的 msg
var msgProcessor MsgProcessor
var jsonResponseProcessor JsonResponseProcessor

// SetMsgProcessor
/*
可以通过此方法实现修改 JsonResponse.Message，可用于: 告知前端此响应是哪个服务返回的，便于后续定位问题.
*/
func SetMsgProcessor(processor MsgProcessor) {
	msgProcessor = processor
}

func ClearMsgProcessor() {
	msgProcessor = nil
}
