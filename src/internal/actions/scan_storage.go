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
    var doResume bool
    if !eutherpeVars.Player.Stopped {
        eutherpeVars.Unlock()
        MusicStop(eutherpeVars, nil)
        eutherpeVars.Lock()
        doResume = true
    }
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
    if eutherpeVars.LoadPlaylists() == nil {
        eutherpeVars.LoadTags()
    }
    eutherpeVars.CollectionHTML = ""
    eutherpeVars.PlaylistsHTML = ""
    if doResume {
        eutherpeVars.Unlock()
        MusicPlay(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    return nil
}
