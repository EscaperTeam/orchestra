package utils

import (
	"os/exec"
)

const Extension = ".exe"

// https://en.wikipedia.org/wiki/Nice_(Unix)
func NiceExec(cmd string, args ...string) *exec.Cmd {
	// Nice not supported on Windows
	return exec.Command(cmd, args...)
}
