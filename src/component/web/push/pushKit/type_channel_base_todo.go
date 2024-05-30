// Package pushKit 此文件中的几个方法，需要在子类中实现!
package pushKit

//func (channel *BaseChannel) BindGroup(group string) {
//	panic("Implement me in the subclass!!!")
//	//BindGroup(channel, group)
//}
//
//func (channel *BaseChannel) BindUser(user string) {
//	panic("Implement me in the subclass!!!")
//	//BindUser(channel, user)
//}
//
//func (channel *BaseChannel) BindBsid(bsid string) {
//	panic("Implement me in the subclass!!!")
//	//BindBsid(channel, bsid)
//}

func (channel *BaseChannel) Initialize() error {
	panic("Implement me in the subclass!!!")
}

func (channel *BaseChannel) Push(data []byte) error {
	panic("Implement me in the subclass!!!")
}

func (channel *BaseChannel) Close(reason string) error {
	panic("Implement me in the subclass!!!")
}
