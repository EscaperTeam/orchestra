// +build !windows

package utils

import (
	"os/exec"
)

const Extension = ""

const niceness = "1"

// https://en.wikipedia.org/wiki/Nice_(Unix)
func NiceExec(cmd string, args ...string) *exec.Cmd {
	return exec.Command("nice", append([]string{"-n", niceness, cmd}, args...)...)
}
