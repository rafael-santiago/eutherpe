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

func PlayModeRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var playModeHTML string
    if len(eutherpeVars.Player.NowPlaying.FilePath) == 0 {
        playModeHTML = "&#x25BA"
    } else {
        playModeHTML = "&#x25A0"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedlePlayMode, playModeHTML, -1)
}
