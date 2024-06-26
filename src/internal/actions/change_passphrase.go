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
)

func ChangePassphrase(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    password, has := (*userData)[vars.EutherpePostFieldPassword]
    if !has {
        return fmt.Errorf("Malformed settings-changepassphrase request.")
    }
    newPassword, has := (*userData)[vars.EutherpePostFieldNewPassword]
    if !has {
        return fmt.Errorf("Malformed settings-changepassphrase request.")
    }
    if len(newPassword[0]) == 0 {
        return fmt.Errorf("Passphrase cannot be null.")
    }
    if !auth.Validate(password[0], eutherpeVars.HTTPd.HashKey) {
        return fmt.Errorf("Wrong passphrase!")
    }
    eutherpeVars.HTTPd.HashKey = auth.HashKey(newPassword[0])
    return nil
}
