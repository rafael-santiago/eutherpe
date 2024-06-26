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

func UpNextRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    upNextHTML := "<ul class=\"nested\">"
    for _, song := range eutherpeVars.Player.UpNext {
        upNextHTML += "<input type=\"checkbox\" id=\"" + song.Artist +
                                                   "/" + song.Album +
                                                   ":" + song.FilePath + "\" class=\"UpNext\">" + song.Title + "<br>"
    }
    upNextHTML += "</ul>"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleUpNext, upNextHTML, -1)
}
