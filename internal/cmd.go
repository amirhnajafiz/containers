//go:build linux
// +build linux

package internal

import (
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/amirhnajafiz/containers/pkg"

	"github.com/google/uuid"
)

// Parent starts a new child process in a new namespace with the specified arguments.
func Parent() error {
	// create a unique container ID
	containerID := uuid.New().String()
	defer func() {
		// clean up the container network after the parent process exits
		if err := removeContainerNetwork(containerID); err != nil {
			// log the error but do not return it, as we are already in a deferred function
			// fmt.Printf("Failed to remove container network: %v\n", err)
		}

		// clean up cgroup after process exits
		if err := removeCgroup(containerID); err != nil {
			// log the error but do not return it, as we are already in a deferred function
			// fmt.Printf("Failed to remove cgroup: %v\n", err)
		}
	}()

	// read configuration settings
	configs, err := pkg.ReadConfigs()
	if err != nil {
		log.Printf("Failed to read configurations: %v\n", err)
	}

	// set up cgroup limits for the container
	memoryLimitMB, _ := strconv.Atoi(configs["memory"])
	cpuShares, _ := strconv.Atoi(configs["cpu"])

	// create cgroup for the container
	if err := createCgroup(containerID, memoryLimitMB, cpuShares); err != nil {
		return err
	}

	// set up the new namespace and execute the child process
	cmd := setContainerNamespace(exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...))
	cmd = setContainerSTD(cmd)

	// start the command in the new namespace
	if err := cmd.Start(); err != nil {
		return err
	}

	// add child PID to cgroup
	if err := addPidToCgroup(containerID, cmd.Process.Pid); err != nil {
		return err
	}

	// create the container network
	if err := setupContainerNetwork(containerID, cmd.Process.Pid, "10.0.0.2"); err != nil {
		return err
	}

	// wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return err
	}

	return nil
}

// Child executes the command specified in the arguments after setting up the container filesystem.
func Child() error {
	// set up the container filesystem
	if err := setContainerFilesystem(); err != nil {
		return err
	}

	// execute the command passed as arguments
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd = setContainerSTD(cmd)

	// run the command in the new namespace
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
