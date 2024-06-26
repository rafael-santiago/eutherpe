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
    "fmt"
    "testing"
)

func TestEutherpeAddrRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.HTTPd.Addr = "127.0.0.1"
    eutherpeVars.HTTPd.Port = "8080"
    templatedInput := fmt.Sprintf("http://%s/index.html", vars.EutherpeTemplateNeedleEutherpeAddr)
    output := EutherpeAddrRender(templatedInput, eutherpeVars)
    if output != "http://127.0.0.1:8080/index.html" {
        t.Errorf("EutherpeAddrRender() seems not to be working accordingly : '%s'.\n", templatedInput)
    }
}
