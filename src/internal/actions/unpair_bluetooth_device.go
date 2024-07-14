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
    "fmt"
    "flag"
)

func UnpairBluetoothDevice(eutherpeVars *vars.EutherpeVars,
                           _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    if len(eutherpeVars.CachedDevices.BlueDevId) == 0 {
        return fmt.Errorf("No device to unpair.")
    }
    _ = bluebraces.DisconnectDevice(eutherpeVars.CachedDevices.BlueDevId, customPath)
    err := bluebraces.UnpairDevice(eutherpeVars.CachedDevices.BlueDevId, customPath)
    if err == nil {
        removeBluetoothDevice(&eutherpeVars.BluetoothDevices, eutherpeVars.CachedDevices.BlueDevId)
        eutherpeVars.CachedDevices.BlueDevId = ""
        eutherpeVars.CachedDevices.MixerControlName = ""
        if len(customPath) == 0 {
            // INFO(Rafael): Since user has chosen unpair the device let's save it asap. It can be
            //               meaningful when user is wanting to power-off her/his device
            //               but wants to unpair the bluetooth first. Withoout saving the session
            //               here, when power-on back again the device will try to pair with the
            //               recently unpaired bluetooth device. Because in normal situations the
            //               session is saved at each 42 secs. I believe that unpair is a special
            //               case that the decision must be saved just after to forwadly reflect
            //               the user intentions accordingly.
            eutherpeVars.SaveSession()
        }
    }
    return err
}

func removeBluetoothDevice(blueDevs *[]bluebraces.BluetoothDevice, blueDevId string) {
    var bIndex int = -1
    for b, currBlueDev := range (*blueDevs) {
        if currBlueDev.Id == blueDevId {
            bIndex = b
            break
        }
    }
    if bIndex == -1 {
        return
    }
    (*blueDevs) = append((*blueDevs)[:bIndex], (*blueDevs)[bIndex + 1:]...)
}
