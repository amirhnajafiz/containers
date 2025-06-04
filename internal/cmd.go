//go:build linux
// +build linux

package internal

import (
	"os"
	"os/exec"
)

// Parent starts a new child process in a new namespace with the specified arguments.
func Parent() error {
	containerID := "mycontainer"
	memoryLimitMB := 128
	cpuShares := 512

	if err := createCgroup(containerID, memoryLimitMB, cpuShares); err != nil {
		return err
	}

	cmd := setContainerNamespace(exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...))
	cmd = setContainerSTD(cmd)

	if err := cmd.Start(); err != nil {
		return err
	}

	// Add child PID to cgroup
	if err := addPidToCgroup(containerID, cmd.Process.Pid); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	// Clean up cgroup after process exits
	if err := removeCgroup(containerID); err != nil {
		return err
	}

	return nil
}

// Child performs the pivot root operation and executes the specified command in the new namespace.
// It mounts the root filesystem, creates an old root directory, and changes the current working directory.
func Child() error {
	if err := setContainerFilesystem(); err != nil {
		return err
	}

	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd = setContainerSTD(cmd)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
