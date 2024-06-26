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

func TestVolumeLevelRender(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Player.VolumeLevel = 42
    output := VolumeLevelRender(vars.EutherpeTemplateNeedleVolumeLevel, eutherpeVars)
    if output != "42" {
        t.Errorf("VolumeLevelRender() is not rendering accordingly : '%s'\n", output)
    }
}
