package actions

import (
    "internal/vars"
    "net/url"
)

func MusicRepeatAll(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    eutherpeVars.Player.RepeatAll = !eutherpeVars.Player.RepeatAll
    eutherpeVars.Player.RepeatOne = false
    return nil
}

func MusicRepeatOne(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    eutherpeVars.Player.RepeatOne = !eutherpeVars.Player.RepeatOne
    eutherpeVars.Player.RepeatAll = false
    return nil
}