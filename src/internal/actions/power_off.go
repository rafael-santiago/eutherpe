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
