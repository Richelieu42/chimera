// Package pushKit 此文件中的几个方法，需要在子类中实现!
package pushKit

func (channel *BaseChannel) Initialize() error {
	panic("Implement me in the subclass!!!")
}

func (channel *BaseChannel) Push(data []byte) error {
	panic("Implement me in the subclass!!!")
}

func (channel *BaseChannel) Close(reason string) error {
	panic("Implement me in the subclass!!!")
}
