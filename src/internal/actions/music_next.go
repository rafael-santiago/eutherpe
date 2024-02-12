package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func MusicNext(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    if eutherpeVars.Player.Stopped ||
       eutherpeVars.Player.Handle == nil  {
        eutherpeVars.Unlock()
        return fmt.Errorf("Not playing anything by now.")
    }
    if eutherpeVars.Player.UpNextCurrentOffset >= len(eutherpeVars.Player.UpNext) - 1 {
        eutherpeVars.Player.UpNextCurrentOffset = -1
    }
    eutherpeVars.Unlock()
    err := MusicStop(eutherpeVars, nil)
    if err != nil {
        return err
    }
    eutherpeVars.Lock()
    eutherpeVars.Player.UpNextCurrentOffset++
    eutherpeVars.Unlock()
    return MusicPlay(eutherpeVars, nil)
}
