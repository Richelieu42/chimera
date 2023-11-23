package cmdKit

import (
	"os/exec"
)

//// NewCommand
///*
//@param setpgid	Windows环境不支持
//@param deathSig	Windows环境不支持
//*/
//func NewCommand(setpgid bool, deathSig syscall.Signal, name string, args ...string) *exec.Cmd {
//	cmd := exec.Command(name, args...)
//
//	//cmd.SysProcAttr = &syscall.SysProcAttr{
//	//	Setpgid:   setpgid,
//	//	Pdeathsig: deathSig,
//	//}
//	return cmd
//}

func (opts CmdOptions) NewCommand(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)

	//cmd.SysProcAttr = &syscall.SysProcAttr{
	//	Setpgid:   opts.Setpgid,
	//	Pdeathsig: opts.Pdeathsig,
	//}
	return cmd
}
