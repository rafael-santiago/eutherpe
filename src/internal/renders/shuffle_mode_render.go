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
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleShuffleMode, shuffleModeHTML, -1)
}
