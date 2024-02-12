package renders

import (
    "internal/vars"
    "encoding/base64"
    "strings"
)

func EncodeAlbumCover(albumCoverBlob string) string {
    return "data:image/" + getImageFmt(albumCoverBlob) + ";base64," +
                        base64.StdEncoding.EncodeToString([]byte(albumCoverBlob))
}

func AlbumArtThumbnailRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var albumArtThumbnailHTML string
    if len(eutherpeVars.Player.NowPlaying.AlbumCover) > 0 {
        albumArtThumbnailHTML = "<img id=\"albumCover\" src=\"" + EncodeAlbumCover(eutherpeVars.Player.NowPlaying.AlbumCover) +
                                    "\" width=125 height=125>"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleAlbumArtThumbnail, albumArtThumbnailHTML, -1)
}

func getImageFmt(blob string) string {
    if strings.HasPrefix(blob, "\x89PNG\r\n\x1A\n") {
        return "png"
    }
    if strings.HasPrefix(blob, "\xFF\xD8") &&
       strings.HasSuffix(blob, "\xFF\xD9") {
        return "jpeg"
    }
    if strings.HasPrefix(blob, "GIF87a") ||
       strings.HasPrefix(blob, "GIF89a") {
        return "gif"
    }
    return "umblauqualquerquenaovaiabrir"
}
