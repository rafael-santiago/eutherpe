//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "testing"
)

func TestAuthenticationModeSwitchRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    output := AuthenticationModeSwitchRender(vars.EutherpeTemplateNeedleAuthenticationModeSwitch, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"flickAuthenticationModeSwitch();\"> Ask passphrase" {
        t.Errorf("AuthenticationModeSwitchRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.HTTPd.Authenticated = true
    output = AuthenticationModeSwitchRender(vars.EutherpeTemplateNeedleAuthenticationModeSwitch, eutherpeVars)
    if output != "<input type=\"checkbox\" onclick=\"flickAuthenticationModeSwitch();\" checked> Ask passphrase" {
        t.Errorf("AuthenticationModeSwitchRender() seems not to be rendering accordingly.\n")
    }
}
