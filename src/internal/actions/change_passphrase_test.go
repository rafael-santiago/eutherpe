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
    "testing"
    "time"
)

func TestChangePassphrase(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.HTTPd.HashKey = auth.HashKey("123mudar*")
    userData := &url.Values{}
    err := ChangePassphrase(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ChangePassphrase() has succeeded when it should fail.\n")
    } else if err.Error() != "Malformed settings-changepassphrase request." {
        t.Errorf("ChangePassphrase() has failed with unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPassword, "123blau")
    err = ChangePassphrase(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ChangePassphrase() has succeeded when it should fail.\n")
    } else if err.Error() != "Malformed settings-changepassphrase request." {
        t.Errorf("ChangePassphrase() has failed with unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldNewPassword, "")
    err = ChangePassphrase(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ChangePassphrase() has succeeded when it should fail.\n")
    } else if err.Error() != "Passphrase cannot be null." {
        t.Errorf("ChangePassphrase() has failed with unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldNewPassword)
    userData.Add(vars.EutherpePostFieldNewPassword, "123mudou")
    err = ChangePassphrase(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ChangePassphrase() has succeeded when it should fail.\n")
    } else if err.Error() != "Wrong passphrase!" {
        t.Errorf("ChangePassphrase() has failed with unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldPassword)
    userData.Add(vars.EutherpePostFieldPassword, "123mudar*")
    err = ChangePassphrase(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ChangePassphrase() has failed when it should succeed.\n")
    }
    time.Sleep(5 * time.Second)
    if !auth.Validate("123mudou", eutherpeVars.HTTPd.HashKey) {
        t.Errorf("ChangePassphrase() seems not to be updating the hash key material.\n")
    }
}
