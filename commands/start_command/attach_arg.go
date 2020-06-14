// +build !windows

package start_command

import (
	"os/exec"
	"syscall"
)

func HookAttachArg(cmd *exec.Cmd) {
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
}
