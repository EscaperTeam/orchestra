package start_command

import "os/exec"

func HookAttachArg(_ *exec.Cmd) {
	// unsupported for windows
}
