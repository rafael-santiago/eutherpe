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
)

func MusicClearAll(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if len(eutherpeVars.Player.UpNext) == 0 {
        return nil
    }
    if !eutherpeVars.Player.Stopped {
        eutherpeVars.Unlock()
        MusicStop(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    if eutherpeVars.Player.Shuffle {
        eutherpeVars.Player.UpNextBkp = make([]mplayer.SongInfo, 0)
        eutherpeVars.Player.Shuffle = false
    }
    eutherpeVars.Player.UpNext = make([]mplayer.SongInfo, 0)
    eutherpeVars.Player.UpNextCurrentOffset = -1
    return nil
}
