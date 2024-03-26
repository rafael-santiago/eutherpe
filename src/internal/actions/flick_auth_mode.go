package actions

import (
    "internal/vars"
    "net/url"
)

func FlickAuthMode(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    eutherpeVars.HTTPd.Authenticated = !eutherpeVars.HTTPd.Authenticated
    return nil
}
