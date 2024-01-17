package storage

import (
    "strings"
    "os"
)

func GetAllAvailableStorages(customPath ...string) []string {
    var availStorages []string
    data, err := os.ReadFile(getMountInfoPath(customPath...))
    if err != nil {
        return availStorages
    }
    lines := strings.Split(string(data), "\n")
    for _, line := range lines {
        fields := strings.Split(line, " ")
        if len(fields) < 10 {
            continue
        }
        if !strings.HasPrefix(fields[9], "/dev/sd") {
            continue
        }
        availStorages = append(availStorages, fields[4])
    }
    return availStorages
}

func getMountInfoPath(customPath ...string) string {
    if len(customPath) > 0 {
        return customPath[0]
    }
    return "/proc/self/mountinfo"
}
