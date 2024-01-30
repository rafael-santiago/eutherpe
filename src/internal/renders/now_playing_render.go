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
