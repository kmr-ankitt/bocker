package cgroups

import (
    "bufio"
    "os"
    "strings"
)

func isCgroupV2() bool {
    f, err := os.Open("/proc/filesystems")
    if err != nil {
        return false
    }
    defer f.Close()

    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        if strings.Contains(scanner.Text(), "cgroup2") {
            return true
        }
    }
    return false
}

func Apply(containerID string, pidLimit, memLimit int) error {
    if isCgroupV2() {
        return ApplyV2(containerID, pidLimit, memLimit)
    }
    return ApplyV1(containerID, pidLimit, memLimit)
}
