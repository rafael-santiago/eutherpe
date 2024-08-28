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
    "flag"
    "net/url"
    "time"
)

func ProbeBluetoothDevices(eutherpeVars *vars.EutherpeVars,
                           _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    blueDevs, err := bluebraces.ScanDevices(time.Duration(10 * time.Second), customPath)
    if err != nil {
        return err
    }
    eutherpeVars.BluetoothDevices = blueDevs
    if len(eutherpeVars.CachedDevices.BlueDevId) == 0 {
        pairedDevices := bluebraces.GetPairedDevices()
        pairedDevicesLen := len(pairedDevices)
        if pairedDevicesLen > 1 {
            unpairNosyDevices(pairedDevices)
        } else if pairedDevicesLen == 1 {
            eutherpeVars.CachedDevices.BlueDevId = pairedDevices[0].Id
        }
    }
    return nil
}

func unpairNosyDevices(pairedDevices []bluebraces.BluetoothDevice) {
    for _, nosyDevice := range pairedDevices {
        bluebraces.DisconnectDevice(nosyDevice.Id)
        bluebraces.UnpairDevice(nosyDevice.Id)
    }
}
