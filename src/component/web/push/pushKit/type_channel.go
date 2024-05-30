package pushKit

type Channel interface {
	// Close 后端主动关闭通道.
	/*
		@param reason 关闭的原因
	*/
	Close(reason string) error
	IsClosed() (rst bool)
	// SetClosed
	/*
		PS: 返回值如果为true，应当 向closeCh发送数据 || 调用Listeners.OnClose().
	*/
	SetClosed() (flag bool)
	// GetCloseCh 获取关闭通道.
	/*
		PS: 只有 调用SetClosed() && 返回值为true 的情况下，才能向通道中写数据.
	*/
	GetCloseCh() chan string

	// Initialize 初始化Channel.
	Initialize() error

	// Dispose 释放Channel所持有的资源.
	Dispose()

	Push(data []byte) error

	GetClientIP() string
	GetType() string

	GetId() string
	GetBsid() string
	SetBsid(string)
	ClearBsid()
	GetUser() string
	SetUser(string)
	ClearUser()
	GetGroup() string
	SetGroup(string)
	ClearGroup()
	GetData() interface{}
	SetData(interface{})
	ClearData()

	//BindGroup(group string)
	//BindUser(user string)
	//BindBsid(bsid string)

	Equals(c Channel) bool
}
