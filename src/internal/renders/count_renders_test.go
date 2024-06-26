//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "internal/mplayer"
    "internal/bluebraces"
    "testing"
)

func TestUpNextCountRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := UpNextCountRender(vars.EutherpeTemplateNeedleUpNextCount,
                                eutherpeVars)
    if output != "0" {
        t.Errorf("UpNextCountRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.Player.UpNext = make([]mplayer.SongInfo, 0)
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, mplayer.SongInfo{})
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, mplayer.SongInfo{})
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, mplayer.SongInfo{})
    output = UpNextCountRender(vars.EutherpeTemplateNeedleUpNextCount,
                               eutherpeVars)
    if output != "3" {
        t.Errorf("UpNextCountRender() seems not to be rendering accordingly.\n")
    }
}

func TestFoundStorageDevicesCountRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := FoundStorageDevicesCountRender(vars.EutherpeTemplateNeedleFoundStorageDevicesCount,
                                             eutherpeVars)
    if output != "0" {
        t.Errorf("FoundStorageDevicesRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.StorageDevices = make([]string, 0)
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "1")
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "2")
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "3")
    output = FoundStorageDevicesCountRender(vars.EutherpeTemplateNeedleFoundStorageDevicesCount,
                                            eutherpeVars)
    if output != "3" {
        t.Errorf("FoundStorageDevicesRender() seems not to be rendering accordingly.\n")
    }
}

func TestFoundBluetoothDevicesCountRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := FoundBluetoothDevicesCountRender(vars.EutherpeTemplateNeedleFoundBluetoothDevicesCount,
                                               eutherpeVars)
    if output != "0" {
        t.Errorf("FoundBluetoothDevicesRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.BluetoothDevices = make([]bluebraces.BluetoothDevice, 0)
    eutherpeVars.BluetoothDevices = append(eutherpeVars.BluetoothDevices, bluebraces.BluetoothDevice{})
    eutherpeVars.BluetoothDevices = append(eutherpeVars.BluetoothDevices, bluebraces.BluetoothDevice{})
    eutherpeVars.BluetoothDevices = append(eutherpeVars.BluetoothDevices, bluebraces.BluetoothDevice{})
    output = FoundBluetoothDevicesCountRender(vars.EutherpeTemplateNeedleFoundBluetoothDevicesCount,
                                              eutherpeVars)
    if output != "3" {
        t.Errorf("FoundBluetoothDevicesRender() seems not to be rendering accordingly.\n")
    }
}
