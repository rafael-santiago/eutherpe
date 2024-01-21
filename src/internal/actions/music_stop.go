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
        time.Sleep(1 * time.Second)
        eutherpeVars.Player.NowPlaying = mplayer.SongInfo{}
    }
    return nil
}
