package renders

import (
    "internal/vars"
    "strings"
)

func PlayerStatusRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var playerStatusJSON string
    nowPlayingInfo := NowPlayingRender(vars.EutherpeTemplateNeedleNowPlaying, eutherpeVars)
    var templatedMarkeeInnerHTML string
    albumCoverSrc := "data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///ywAAAAAAQABAAACAUwAOw=="
    if len(nowPlayingInfo) > 0 {
        templatedMarkeeInnerHTML = nowPlayingInfo
        albumCoverSrc = EncodeAlbumCover(eutherpeVars.Player.NowPlaying.AlbumCover)
    }
    playerStatusJSON = "{\"now-playing\":\"" + templatedMarkeeInnerHTML + "\"," +
                        "\"album-cover-src\" : \"" + albumCoverSrc + "\"}"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedlePlayerStatus, playerStatusJSON, -1)
}
