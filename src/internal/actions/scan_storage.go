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
    "internal/mplayer"
    "net/url"
    "fmt"
)

func ScanStorage(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return fmt.Errorf("Unset MusicDevId.")
    }
    newCollection, err := mplayer.LoadMusicCollection(eutherpeVars.CachedDevices.MusicDevId, eutherpeVars.GetCoversCacheRootPath())
    if err != nil {
        return err
    }
    if len(eutherpeVars.Collection) > 0 {
        eutherpeVars.SaveCollection()
    }
    eutherpeVars.Collection = newCollection
    if len(eutherpeVars.Collection) > 0 {
        eutherpeVars.SaveCollection()
    }
    return nil
}
