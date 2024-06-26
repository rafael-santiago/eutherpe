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

func TestLastErrorRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleLastError)
    output := LastErrorRender(templatedInput, eutherpeVars)
    if len(output) != 0 {
        t.Errorf("LastErrorRender() seems not to be rendering accordingly.\n")
    }
    eutherpeVars.LastError = fmt.Errorf("Deu Merda! Vc Nao Eh Quadrado! Se vira!")
    output = LastErrorRender(templatedInput, eutherpeVars)
    if output != "Deu Merda! Vc Nao Eh Quadrado! Se vira!" {
        t.Errorf("LastErrorRender() seems not to be rendering accordingly.\n")
    }
}
