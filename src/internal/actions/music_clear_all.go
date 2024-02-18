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
    eutherpeVars.Player.UpNext = make([]mplayer.SongInfo, 0)
    return nil
}
