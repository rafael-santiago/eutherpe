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

func TestCurrentConfigRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.CurrentConfig = "UnitTest"
    output := CurrentConfigRender(vars.EutherpeTemplateNeedleCurrentConfig, eutherpeVars)
    if output != "UnitTest" {
        t.Errorf("CurrentConfigRender() is not rendering accordingly.\n")
    }
}
