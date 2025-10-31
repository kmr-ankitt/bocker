package cgroups

import (
    "fmt"
    "os"
    "path/filepath"
)

const cgroupRootV1 = "/sys/fs/cgroup"

func ApplyV1(containerID string, pidLimit, memLimit int) error {
    // --- PID cgroup ---
    pidPath := filepath.Join(cgroupRootV1, "pids", containerID)

    os.MkdirAll(pidPath, 0755)

    os.WriteFile(filepath.Join(pidPath, "pids.max"), []byte(fmt.Sprintf("%d", pidLimit)), 0644)

    os.WriteFile(filepath.Join(pidPath, "notify_on_release"), []byte("1"), 0644)

    os.WriteFile(filepath.Join(pidPath, "cgroup.procs"), []byte(fmt.Sprintf("%d", os.Getpid())), 0644)

    // --- Memory cgroup ---
    memPath := filepath.Join(cgroupRootV1, "memory", containerID)

    os.MkdirAll(memPath, 0755)

    os.WriteFile(filepath.Join(memPath, "memory.limit_in_bytes"), []byte(fmt.Sprintf("%d", memLimit*1024*1024)), 0644)

    os.WriteFile(filepath.Join(memPath, "cgroup.procs"), []byte(fmt.Sprintf("%d", os.Getpid())), 0644)

    return nil
}
