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
    "internal/dj"
    "strings"
)

func PlaylistsRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    if len(eutherpeVars.PlaylistsHTML) == 0 {
        eutherpeVars.PlaylistsHTML = renderPlaylists(eutherpeVars.Playlists)
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedlePlaylists,
                           eutherpeVars.PlaylistsHTML, 1)
}

func renderPlaylists(playlists []dj.Playlist) string {
    var playlistsHTML string
    for _, playlist := range playlists {
        playlistsHTML += "<ul id=\"eutherpeUL\"><li>"
        playlistName := playlist.Name
        playlistsHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);selectPlaylist(this);\" id=\"" + playlistName + "\" class=\"Playlist\"><span class=\"caret\">" + playlistName + "</span>"
        playlistsHTML += "<ul class=\"nested\">"
        for _, song := range playlist.Songs() {
            playlistsHTML += "<li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"" + playlistName + ":" + song.Artist + "/" + song.Album + "/:" + song.FilePath + "\" class=\"PlaylistSong\">" + song.Title + "</li>"
        }
        playlistsHTML += "</ul>"
        playlistsHTML += "</li></ul>"
    }
    return playlistsHTML
}