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
    "time"
)

func MusicStop(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if eutherpeVars.Player.Stopped {
        return nil
    }
    if eutherpeVars.Player.Handle != nil {
        eutherpeVars.Player.Stopped = true
        mplayer.Stop(eutherpeVars.Player.Handle)
        eutherpeVars.Player.Handle = nil
        time.Sleep(10 * time.Nanosecond)
        eutherpeVars.Player.NowPlaying = mplayer.SongInfo{}
    }
    return nil
}
