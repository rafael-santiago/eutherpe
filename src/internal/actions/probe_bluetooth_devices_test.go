//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package actions

import (
    "internal/vars"
    "internal/bluebraces"
    "net/url"
    "os"
    "testing"
)

func TestProbeBluetoothDevices(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.CachedDevices.BlueDevId = "/dev/blue-42"
    userData := &url.Values{}
    err := ProbeBluetoothDevices(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ProbeBluetoothDevices() has failed when it should not.\n")
    }
    expected := []bluebraces.BluetoothDevice {
        {"E3:91:B6:02:8C:47", "GT FUN"},
        {"B5:D0:38:C0:ED:74", "EASYWAY-BLE"},
        {"BA:BA:CA:BA:BA:CA", "PHONE-BLAU"},
        {"42:42:42:42:42:42", "zaphoid-spks"},
    }
    if eutherpeVars.CachedDevices.BlueDevId != "/dev/blue-42" {
        t.Errorf("ProbeBluetoothDevices() seems to be clearing the cached bluetooth device id.\n")
    }
    if len(eutherpeVars.BluetoothDevices) != len(expected) {
        t.Errorf("Wrong total of bluetooth devices.\n")
    } else {
        for d, _ := range expected {
            if eutherpeVars.BluetoothDevices[d] != expected[d] {
                t.Errorf("{ '%v' != '%v' }\n", eutherpeVars.BluetoothDevices[d], expected[d])
            }
        }
    }
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    wtfDev := bluebraces.BluetoothDevice { "FF:FF:FF:FF:FF:FF", "WTF DEV" }
    eutherpeVars.BluetoothDevices = []bluebraces.BluetoothDevice {
        wtfDev,
    }
    err = ProbeBluetoothDevices(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ProbeBluetoothDevices() has not failed when it should.\n")
    } else if err.Error() != "exit status 1" {
        t.Errorf("ProbeBluetoothDevices() has returned unexpected error : '%v'.\n", err.Error())
    }
    if len(eutherpeVars.BluetoothDevices) != 1 {
        t.Errorf("ProbeBluetoothDevices() seems to be changing []BluetoothDevices on error cases.\n")
    } else if eutherpeVars.BluetoothDevices[0] != wtfDev {
        t.Errorf("ProbeBluetoothDevices() seems to be corrupting []BluetoothDevices on error cases.\n")
    }
}
