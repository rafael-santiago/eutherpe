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
    "net/url"
    "fmt"
)

func SetCurrentConfig(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var err error
    currentConfig := userData.Get(vars.EutherpePostFieldConfig)
    switch currentConfig {
        case vars.EutherpeConfigMusic,
             vars.EutherpeConfigCollection,
             vars.EutherpeConfigPlaylists,
             vars.EutherpeConfigStorage,
             vars.EutherpeConfigBluetooth,
             vars.EutherpeConfigSettings:
                eutherpeVars.CurrentConfig = currentConfig
                break
        default:
            err = fmt.Errorf("Unknown config value.")
            break
    }
    return err
}
