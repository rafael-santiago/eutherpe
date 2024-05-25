package actions

import (
    "internal/vars"
    "internal/system"
    "net/url"
)

func PowerOff(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if !eutherpeVars.Player.Stopped {
        eutherpeVars.Unlock()
        MusicStop(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    return system.Shutdown()
}
