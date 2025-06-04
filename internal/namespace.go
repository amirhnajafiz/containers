//go:build linux
// +build linux

package internal

import (
	"os/exec"
	"syscall"
)

// setContainerNamespace sets the necessary namespace flags for the command to run in a new container namespace.
func setContainerNamespace(cmd *exec.Cmd) *exec.Cmd {
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	return cmd
}
