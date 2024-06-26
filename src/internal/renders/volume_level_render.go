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
    "strings"
    "fmt"
)

func VolumeLevelRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    volumeLevel := fmt.Sprintf("%d", eutherpeVars.Player.VolumeLevel)
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleVolumeLevel, volumeLevel, -1)
}
