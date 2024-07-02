//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package bluebraces

import(
    "os"
    "os/exec"
    "time"
    "strings"
    "path"
)

const (
    kNextDeviceNeedle = "] Device "
    kDevIdLen = 17
    kAliasNeedle = "Alias: "
)

var g_HasSystemCtlCallNr int
var g_HasSystemCtl bool

func GetPairedDevices(customPath ...string) []BluetoothDevice {
    cp := getToolPath(customPath...)
    out, err := exec.Command(cp + "bluetoothctl", "paired-devices").CombinedOutput()
    pairedDevices := make([]BluetoothDevice, 0)
    if err != nil {
        return pairedDevices
    }
    foundDevices := strings.Split(string(out), "\n")
    for _, device := range foundDevices {
        off := strings.Index(device, "Device ")
        if off == -1 {
            continue
        }
        off = strings.Index(device[7:], " ")
        if off == -1 {
            continue
        }
        off += 7
        newDevice := BluetoothDevice{}
        newDevice.Id = device[7:off]
        if len(device) < off + 1 {
            continue
        }
        newDevice.Alias = device[off+1:]
        pairedDevices = append(pairedDevices, newDevice)
    }
    return pairedDevices
}

func Wear(customPath ...string) error {
    cp := getToolPath(customPath...)
    err := exec.Command(cp + "bluetoothctl", "power", "on").Run()
    if err == nil {
        err = exec.Command(cp + "bluetoothctl", "discoverable", "on").Run()
    }
    return err
}

func Unwear(customPath ...string) error {
    cp := getToolPath(customPath...)
    err := exec.Command(cp + "bluetoothctl", "power", "off").Run()
    if err != nil {
        return err
    }
    return err
}

func ScanDevices(duration time.Duration, customPath ...string) ([]BluetoothDevice, error) {
    blueDevs := make([]BluetoothDevice, 0)
    err := doDevicesScan(&blueDevs, duration, customPath...)
    return blueDevs, err
}

func PairDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "pair", devId).Run()
}

func UnpairDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "remove", devId).Run()
}

func ConnectDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "connect", devId).Run()
}

func DisconnectDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "disconnect", devId).Run()
}

func TrustDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "trust", devId).Run()
}

func UntrustDevice(devId string, customPath ...string) error {
    return exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "untrust", devId).Run()
}

func getToolPath(customPath ...string) string {
    if len(customPath) > 0 {
        return customPath[0]
    }
    return ""
}

func doDevicesScan(blueDevs *[]BluetoothDevice, duration time.Duration, customPath ...string) error {
    cmd := exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "scan", "on")
    var out []byte
    var err error
    go func() {
        out, err = cmd.CombinedOutput()
    }()

    time.Sleep(duration)
    cmd.Process.Signal(os.Interrupt)
    time.Sleep(time.Duration(1 * time.Second))

    if err != nil {
        return err
    }

    sOut := string(out)
    nextDevOff := strings.Index(sOut, kNextDeviceNeedle)
    var startOff int
    for nextDevOff > -1 {
        startOff += nextDevOff + len(kNextDeviceNeedle)
        id := sOut[startOff : startOff + kDevIdLen]
        if !hasDeviceById(id, *blueDevs) {
            *blueDevs = append(*blueDevs, BluetoothDevice { id, getDeviceAlias(id, customPath...) })
        }
        nextDevOff = strings.Index(sOut[startOff:], kNextDeviceNeedle)
    }

    return nil
}

func hasDeviceById(id string, blueDevs []BluetoothDevice) bool {
    if blueDevs == nil {
        return false
    }
    for _, blueDev := range blueDevs {
        if blueDev.Id == id {
            return true
        }
    }
    return false
}

func getDeviceAlias(devId string, customPath ...string) string {
    cmd := exec.Command(path.Join(getToolPath(customPath...), "bluetoothctl"), "info", devId)
    out, err := cmd.CombinedOutput()
    if err != nil {
        return "(unamed)"
    }
    sOut := string(out)
    aliasOff := strings.Index(sOut, kAliasNeedle)
    if aliasOff < 0 {
        return "(unamed)"
    }
    var alias string
    for a := aliasOff + len(kAliasNeedle); a < len(sOut) && sOut[a] != '\n' && sOut[a] != '\r'; a++ {
        alias += string(sOut[a])
    }
    return alias
}

func hasSystemCtl() bool {
    if g_HasSystemCtlCallNr == 0 {
        err := exec.Command("systemctl", "status").Run()
        g_HasSystemCtl = (err == nil)
        g_HasSystemCtlCallNr++
    }
    return g_HasSystemCtl
}
