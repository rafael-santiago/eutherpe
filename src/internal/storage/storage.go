package storage

import (
    "strings"
    "os"
    "os/exec"
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

func GetDeviceSerialNumberByMountPoint(mountPoint string, customPath ...string) string {
    devPath := getDeviceByMountPoint(mountPoint, customPath...)
    if len(devPath) == 0 {
        return ""
    }
    outLines := getUSBInfo(devPath)
    if len(outLines) == 0 {
        return ""
    }
    serial := getUdevadmInfoData("ID_USB_SERIAL_SHORT", outLines)
    if len(serial) > 0 {
        return serial
    }
    serial = getUdevadmInfoData("ID_SERIAL_SHORT", outLines)
    if len(serial) > 0 {
        return serial
    }
    serial = getUdevadmInfoData("ID_VENDOR_ID", outLines)
    if len(serial) == 0 {
        return ""
    }
    serial += getUdevadmInfoData("ID_MODEL_ID", outLines)
    return serial
}

func getDeviceByMountPoint(mountPoint string, customPath ...string) string {
    data, err := os.ReadFile(getMountInfoPath(customPath...))
    if err != nil {
        return ""
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
        if fields[4] == mountPoint {
            return fields[9]
        }
    }
    return ""
}

func getUSBInfo(devPath string, customPath ...string) []string {
    udevadmPath := "udevadm"
    if len(customPath) > 0 {
        udevadmPath = customPath[0] + udevadmPath
    }
    out, err := exec.Command(udevadmPath, "info", "--name=" + devPath).CombinedOutput()
    if err != nil {
        return make([]string, 0)
    }
    return strings.Split(string(out), "\n")
}

func getUdevadmInfoData(field string, outLines []string) string {
    needle := field + "="
    for _, line := range outLines {
        l := strings.Index(line, needle)
        if l != -1 {
            return line[l + len(needle):]
        }
    }
    return ""
}
