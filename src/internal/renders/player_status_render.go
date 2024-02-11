package renders

import (
    "internal/vars"
    "strings"
)

func PlayerStatusRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var playerStatusJSON string
    templatedMarkeeInnerHTML := "&#119066;&#119070;&#119066;&#119046; " +
                                    vars.EutherpeTemplateNeedleNowPlaying +
                                        " &#119066;&#119047;"
    playerStatusJSON = "{\"now-playing-markee\":\"" +
                            NowPlayingRender(templatedMarkeeInnerHTML, eutherpeVars) +  "\"}"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedlePlayerStatus, playerStatusJSON, -1)
}
