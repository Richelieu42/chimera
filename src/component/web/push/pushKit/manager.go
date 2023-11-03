package pushKit

import (
	"github.com/richelieu-yang/chimera/v2/src/component/web/push/pushKit/types"
	"github.com/richelieu-yang/chimera/v2/src/core/mapKit"
	"github.com/richelieu-yang/chimera/v2/src/core/setKit"
)

var (
	// idMap 即allMap
	/*
		key: id（一对一）
	*/
	idMap = mapKit.NewMapWithLock[string, types.Channel]()

	// bsidMap
	/*
		key: bsid（一对一）
	*/
	bsidMap = mapKit.NewMapWithLock[string, types.Channel]()

	// userMap
	/*
		key: user（一对多）
	*/
	userMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[types.Channel]]()

	// groupMap
	/*
		key: group（一对多）
	*/
	groupMap = mapKit.NewMapWithLock[string, *setKit.SetWithLock[types.Channel]]()
)

// GetChannelByBsid （读锁）
/*
@return 可能为nil
*/
func GetChannelByBsid(bsid string) (channel types.Channel) {
	/* 读锁 */
	bsidMap.RWLock.RLockFunc(func() {
		channel = bsidMap.Map[bsid]
	})
	return
}

// GetUserSet （读锁）
/*
@return 可能为nil
*/
func GetUserSet(user string) *setKit.SetWithLock[types.Channel] {
	var userSet *setKit.SetWithLock[types.Channel]
	/* 读锁 */
	userMap.RWLock.RLockFunc(func() {
		userSet = userMap.Map[user]
	})
	return userSet
}

// GetGroupSet （读锁）
/*
@return 可能为nil
*/
func GetGroupSet(group string) *setKit.SetWithLock[types.Channel] {
	var groupSet *setKit.SetWithLock[types.Channel]
	/* 读锁 */
	groupMap.RWLock.RLockFunc(func() {
		groupSet = userMap.Map[group]
	})
	return groupSet
}
