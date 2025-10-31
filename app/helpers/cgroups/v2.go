package cgroups

import (
    "fmt"
    "os"
    "path/filepath"
)

const cgroupRootV2 = "/sys/fs/cgroup"

func ApplyV2(containerID string, pidLimit, memLimit int) error {
    path := filepath.Join(cgroupRootV2, containerID)
    os.MkdirAll(path, 0755)

    // enable controllers
    os.WriteFile(filepath.Join(cgroupRootV2, "cgroup.subtree_control"), []byte("+pids +memory"), 0644)

    // memory
    os.WriteFile(filepath.Join(path, "memory.max"), []byte(fmt.Sprintf("%d", memLimit*1024*1024)), 0644)

    // pids
    os.WriteFile(filepath.Join(path, "pids.max"), []byte(fmt.Sprintf("%d", pidLimit)), 0644)

    // attach this process
    return os.WriteFile(filepath.Join(path, "cgroup.procs"), []byte(fmt.Sprintf("%d", os.Getpid())), 0644)
}
