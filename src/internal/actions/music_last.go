package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func MusicLast(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    if eutherpeVars.Player.Stopped ||
       eutherpeVars.Player.Handle == nil  {
        eutherpeVars.Unlock()
        return fmt.Errorf("Not playing anything by now.")
    }
    if eutherpeVars.Player.UpNextCurrentOffset < 0 {
        eutherpeVars.Player.UpNextCurrentOffset = len(eutherpeVars.Player.UpNext)
    }
    eutherpeVars.Unlock()
    err := MusicStop(eutherpeVars, nil)
    if err != nil {
        return err
    }
    eutherpeVars.Lock()
    eutherpeVars.Player.UpNextCurrentOffset--
    eutherpeVars.Unlock()
    return MusicPlay(eutherpeVars, nil)
}
