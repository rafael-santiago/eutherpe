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
    "internal/storage"
    "net/url"
)

func ListStorages(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    eutherpeVars.StorageDevices = storage.GetAllAvailableStorages()
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return nil
    }
    var found bool
    for _, device := range eutherpeVars.StorageDevices {
        found = (device == eutherpeVars.CachedDevices.MusicDevId)
        if found {
            break
        }
    }
    if !found {
        eutherpeVars.CachedDevices.MusicDevId = ""
    }
    return nil
}
