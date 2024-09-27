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
    "flag"
)

func Convert2MP3(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    var doResume bool
    if !eutherpeVars.Player.Stopped {
        eutherpeVars.Unlock()
        MusicStop(eutherpeVars, nil)
        eutherpeVars.Lock()
        doResume = true
    }
    defer eutherpeVars.Unlock()
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../mplayer"
    }
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return fmt.Errorf("You need to set a storage device first.")
    }
    err := mplayer.ConvertSongs(eutherpeVars.CachedDevices.MusicDevId, customPath)
    if doResume {
        eutherpeVars.Unlock()
        MusicPlay(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    return err
}
