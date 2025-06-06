//go:build linux
// +build linux

package internal

import (
	"fmt"
	"os"
	"syscall"
)

// setContainerFilesystem mounts the root filesystem, creates an old root directory,
// and performs the pivot root operation to switch the root filesystem of the process.
func setContainerFilesystem(workdir string) error {
	if err := syscall.Mount("rootfs", "rootfs", "", syscall.MS_BIND, ""); err != nil {
		return fmt.Errorf("failed to mount rootfs: %w", err)
	}
	if err := os.MkdirAll("rootfs/oldrootfs", 0700); err != nil {
		return fmt.Errorf("failed to create oldrootfs directory: %w", err)
	}
	if err := syscall.PivotRoot("rootfs", "rootfs/oldrootfs"); err != nil {
		return fmt.Errorf("failed to pivot root: %w", err)
	}
	if err := os.Chdir("/"); err != nil {
		return fmt.Errorf("failed to change directory to new root: %w", err)
	}

	// copy the workdir to the new root
	if workdir == "" {
		if err := syscall.Mount(workdir, workdir, "", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
			return fmt.Errorf("failed to mount workdir: %w", err)
		}
		if err := syscall.Chroot("."); err != nil {
			return fmt.Errorf("failed to chroot into new root: %w", err)
		}
	}

	return nil
}
