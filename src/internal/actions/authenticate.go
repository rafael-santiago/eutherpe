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
    "internal/auth"
    "net/url"
    "fmt"
    "strings"
)

func Authenticate(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    action, has := (*userData)[vars.EutherpeActionId]
    if !has || action[0] != vars.EutherpeAuthenticateId {
        return fmt.Errorf("Malformed authenticate request.")
    }
    remoteAddr, has := (*userData)[vars.EutherpePostFieldRemoteAddr]
    if !has {
        return fmt.Errorf("Malformed authenticate request.")
    }
    password, has := (*userData)[vars.EutherpePostFieldPassword]
    if !has {
        return fmt.Errorf("Malformed authenticate request.")
    }
    if !auth.Validate(password[0], eutherpeVars.HTTPd.HashKey) {
        return fmt.Errorf("Wrong passphrase!")
    }
    p := strings.Index(remoteAddr[0], ":")
    if p > -1 {
        remoteAddr[0] = remoteAddr[0][0:p]
    }
    eutherpeVars.HTTPd.AuthWatchdog.RefreshAuthWindow(remoteAddr[0])
    return nil
}