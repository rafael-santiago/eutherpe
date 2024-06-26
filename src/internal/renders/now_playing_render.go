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

func NowPlayingRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var nowPlayingHTML string
    if len(eutherpeVars.Player.NowPlaying.Title) > 0 {
        nowPlayingHTML = eutherpeVars.Player.NowPlaying.Artist + " - " + eutherpeVars.Player.NowPlaying.Title
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleNowPlaying, nowPlayingHTML, -1)
}
