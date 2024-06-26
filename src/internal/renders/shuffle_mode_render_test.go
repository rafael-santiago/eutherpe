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

func TestShuffleModeRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleShuffleMode)
    output := ShuffleModeRender(templatedInput, eutherpeVars)
    if output != "Shuffle" {
        t.Errorf("ShuffleModeRender() seems not to be rendering accordingly.\n")
    } else {
        eutherpeVars.Player.Shuffle = true
        output = ShuffleModeRender(templatedInput, eutherpeVars)
        if output != "Original" {
            t.Errorf("ShuffleModeRender() seems not to be rendering accordingly.\n")
        }
    }
}
