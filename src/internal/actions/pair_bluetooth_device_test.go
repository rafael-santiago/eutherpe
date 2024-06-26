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
    _"internal/bluebraces"
    "net/url"
    "os"
    "testing"
)

func TestPairBluetoothDevice(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := PairBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("PairBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "Malformed bluetooth-pair request." {
        t.Errorf("PairBluetoothDevice() has failed with unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldBluetoothDevice, "Blue42")
    err = PairBluetoothDevice(eutherpeVars, userData)
    if err != nil {
        t.Errorf("PairBluetoothDevice() has failed when it should not.\n")
    } else if eutherpeVars.CachedDevices.BlueDevId != "Blue42" {
        t.Errorf("PairBluetoothDevices() seems not to be caching the device id after a successful pairing.\n")
    }
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    userData.Del(vars.EutherpePostFieldBluetoothDevice)
    userData.Add(vars.EutherpePostFieldBluetoothDevice, "Vogon_sPoetryRandomDev")
    err = PairBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("PairBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "exit status 1" {
        t.Errorf("PairBluetoothDevice() has failed with unexpected error.\n")
    }
    if eutherpeVars.CachedDevices.BlueDevId != "Blue42" {
        t.Errorf("PairBluetoothDevice() seems to be messing with cached device even when failing.\n")
    }
}
