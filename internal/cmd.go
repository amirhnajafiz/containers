//go:build linux
// +build linux

package internal

import (
	"os"
	"os/exec"
	"syscall"
)

// Parent starts a new child process in a new namespace with the specified arguments.
func Parent() error {
	cmd := setContainerNamespace(exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...))
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

// Child performs the pivot root operation and executes the specified command in the new namespace.
// It mounts the root filesystem, creates an old root directory, and changes the current working directory.
func Child() error {
	if err := syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""); err != nil {
		return err
	}
	if err := os.MkdirAll("rootfs/oldrootfs", 0700); err != nil {
		return err
	}
	if err := syscall.PivotRoot("rootfs", "rootfs/oldrootfs"); err != nil {
		return err
	}
	if err := os.Chdir("/"); err != nil {
		return err
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
