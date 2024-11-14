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
    "encoding/base64"
    "strings"
    "os"
    "path"
)

var b64ImageData string

func EncodeAlbumCover(albumCoverBlob string) string {
    var imageFormat string
    if len(b64ImageData) > 0 {
        return b64ImageData
    }
    if len(albumCoverBlob) > 0 {
        imageFormat = getImageFmt(albumCoverBlob)
        b64ImageData = base64.StdEncoding.EncodeToString([]byte(albumCoverBlob))
    } else {
        imageFormat = "png"
        b64ImageData = getUncoveredAlbumImageData()
    }
    b64ImageData = "data:image/" + imageFormat  + ";base64," + b64ImageData
    return b64ImageData
}

func AlbumArtThumbnailRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    albumArtThumbnailHTML := eutherpeVars.RenderedAlbumArtThumbnailHTML
    if !eutherpeVars.Player.Stopped && len(albumArtThumbnailHTML) == 0 {
        b64ImageData = ""
        eutherpeVars.Player.NowPlaying.AlbumCover = getAlbumCoverBlob(eutherpeVars.GetCoversCacheRootPath(), eutherpeVars.Player.NowPlaying.AlbumCover)
        albumArtThumbnailHTML = "<img id=\"albumCover\" src=\"" + EncodeAlbumCover(eutherpeVars.Player.NowPlaying.AlbumCover) +
                                "\" width=125 height=125 onclick=\"openAlbumCoverViewer();\" class=\"AlbumCover\">"
        eutherpeVars.RenderedAlbumArtThumbnailHTML = albumArtThumbnailHTML
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleAlbumArtThumbnail, albumArtThumbnailHTML, 1)
}

func getAlbumCoverBlob(coversCacheRootPath, albumCoverBlob string) string {
    if !strings.HasPrefix(albumCoverBlob, "blob-id=") {
        return albumCoverBlob
    }
    blob, err := os.ReadFile(path.Join(coversCacheRootPath, albumCoverBlob[8:]))
    if err != nil {
        return ""
    }
    return string(blob)
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
    return ""
}

func getUncoveredAlbumImageData() string {
    return "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"
}
