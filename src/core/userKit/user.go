package userKit

import (
	"log"
	"os/user"
	"sync"
)

var (
	uOnce sync.Once
	u     *user.User
)

func getU() *user.User {
	// 使用 sync.Once 实现单例懒加载
	uOnce.Do(func() {
		var err error

		u, err = user.Current()
		if err != nil {
			log.Fatalf("Fail to get current user, error: %s", err.Error())
		}
	})

	return u
}

// GetUid user ID
func GetUid() string {
	return getU().Uid
}

// GetGid primary group ID
func GetGid() string {
	return getU().Gid
}

func GetName() string {
	return getU().Name
}

func GetUserName() string {
	return getU().Username
}

// GetUserHomeDir 获取当前用户的目录.
/*
@return 必定不为"" && 是个存在的目录

e.g.
() => "/Users/richelieu"
*/
func GetUserHomeDir() string {
	return getU().HomeDir
}

//func getUserHomeDir() (string, error) {
//	//// os.Getenv("user.home")可能会返回""，比如在Mac环境下
//	//userHomeDir := os.Getenv("user.home")
//	//if userHomeDir == "" {
//	//	userHomeDir = os.Getenv("HOME")
//	//}
//
//	return homedir.Dir()
//}
