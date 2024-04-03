package actions

import (
    "internal/vars"
    "net/url"
)

func FlickHTTPSMode(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    eutherpeVars.HTTPd.TLS = !eutherpeVars.HTTPd.TLS
    // TODO(Rafael): Generate certificate if it was not found.
    return nil
}
