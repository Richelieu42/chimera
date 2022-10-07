package userKit

import (
	"github.com/mitchellh/go-homedir"
	"github.com/richelieu42/go-scales/src/core/errorKit"
	"os/user"
)

var homeDir string
var u *user.User

func init() {
	var err error

	homeDir, err = homedir.Dir()
	if err != nil {
		errorKit.Panic("homedir.Dir() fails, error:\n%+v", err)
	}

	u, err = user.Current()
	if err != nil {
		errorKit.Panic("user.Current() fails, error:\n%+v", err)
	}
}
