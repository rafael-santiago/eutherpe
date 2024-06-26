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
    "internal/wifi"
    "net/url"
    "fmt"
)

func SetWLANCredentials(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    ESSID, has := (*userData)[vars.EutherpePostFieldESSID]
    if !has {
        return fmt.Errorf("Malformed settings-setwlancredentials request.")
    }
    if len(ESSID[0]) == 0 {
        if eutherpeVars.WLAN.ConnSession != nil {
            wifi.Stop(eutherpeVars.WLAN.ConnSession)
            eutherpeVars.WLAN.ConnSession = nil
        }
        eutherpeVars.WLAN.ESSID = ""
        return nil
    }
    password, has := (*userData)[vars.EutherpePostFieldPassword]
    if !has {
        return fmt.Errorf("Malformed settings-setwlancredentials request.")
    }
    err := wifi.SetWPAPassphrase(ESSID[0], password[0])
    if err ==  nil {
        eutherpeVars.WLAN.ESSID = ESSID[0]
    }
    return err
}
