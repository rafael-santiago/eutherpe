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
