package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// createCgroup sets up a cgroup for the container with memory and CPU limits.
func createCgroup(containerID string, memoryLimitMB int, cpuShares int) error {
	cgroupBase := "/sys/fs/cgroup"
	memoryPath := filepath.Join(cgroupBase, "memory", containerID)
	cpuPath := filepath.Join(cgroupBase, "cpu", containerID)

	// Create memory cgroup
	if err := os.MkdirAll(memoryPath, 0755); err != nil {
		return fmt.Errorf("create memory cgroup: %w", err)
	}
	if err := os.WriteFile(filepath.Join(memoryPath, "memory.limit_in_bytes"), []byte(strconv.Itoa(memoryLimitMB*1024*1024)), 0644); err != nil {
		return fmt.Errorf("set memory limit: %w", err)
	}

	// Create CPU cgroup
	if err := os.MkdirAll(cpuPath, 0755); err != nil {
		return fmt.Errorf("create cpu cgroup: %w", err)
	}
	if err := os.WriteFile(filepath.Join(cpuPath, "cpu.shares"), []byte(strconv.Itoa(cpuShares)), 0644); err != nil {
		return fmt.Errorf("set cpu shares: %w", err)
	}

	return nil
}

// addPidToCgroup adds a process to the cgroup.
func addPidToCgroup(containerID string, pid int) error {
	cgroupBase := "/sys/fs/cgroup"
	memoryTasks := filepath.Join(cgroupBase, "memory", containerID, "tasks")
	cpuTasks := filepath.Join(cgroupBase, "cpu", containerID, "tasks")

	pidStr := []byte(strconv.Itoa(pid))

	if err := os.WriteFile(memoryTasks, pidStr, 0644); err != nil {
		return fmt.Errorf("add pid to memory cgroup: %w", err)
	}
	if err := os.WriteFile(cpuTasks, pidStr, 0644); err != nil {
		return fmt.Errorf("add pid to cpu cgroup: %w", err)
	}

	return nil
}

// removeCgroup cleans up the cgroup after the container exits.
func removeCgroup(containerID string) error {
	cgroupBase := "/sys/fs/cgroup"
	memoryPath := filepath.Join(cgroupBase, "memory", containerID)
	cpuPath := filepath.Join(cgroupBase, "cpu", containerID)

	if err := os.RemoveAll(memoryPath); err != nil {
		return fmt.Errorf("remove memory cgroup: %w", err)
	}
	if err := os.RemoveAll(cpuPath); err != nil {
		return fmt.Errorf("remove cpu cgroup: %w", err)
	}
	return nil
}
