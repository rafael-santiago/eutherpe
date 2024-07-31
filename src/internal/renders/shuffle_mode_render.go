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
)

func ShuffleModeRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var shuffleModeHTML string
    if eutherpeVars.Player.Shuffle {
        shuffleModeHTML = "Original"
    } else {
        shuffleModeHTML = "Shuffle"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleShuffleMode, shuffleModeHTML, 1)
}
