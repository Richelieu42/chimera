package pushKit

func (channel *BaseChannel) Initialize() error {
	//TODO implement me
	panic("implement me")
}

func (channel *BaseChannel) Close(reason string) error {
	//TODO implement me
	panic("implement me")
}

func (channel *BaseChannel) Dispose() {
	channel.Interval.Stop()
	channel.Interval = nil
}

func (channel *BaseChannel) Push(data []byte) error {
	//TODO implement me
	panic("implement me")
}
