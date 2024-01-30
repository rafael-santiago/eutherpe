package renders

import (
    "internal/vars"
    "strings"
)

func PlaylistsRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var playlistsHTML string
    for _, playlist := range eutherpeVars.Playlists {
        playlistsHTML += "<ul id=\"eutherpeUL\"><li>"
        playlistName := playlist.Name
        playlistsHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);selectPlaylist(this);\" id=\"" + playlistName + "\" class=\"Playlist\"><span class=\"caret\">" + playlistName + "</span>"
        playlistsHTML += "<ul class=\"nested\">"
        for _, song := range playlist.Songs() {
            playlistsHTML += "<li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"" + playlistName + ":" + song.FilePath + "\" class=\"PlaylistSong\">" + song.Title + "</li>"
        }
        playlistsHTML += "</ul>"
        playlistsHTML += "</li></ul>"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedlePlaylists, playlistsHTML, -1)
}