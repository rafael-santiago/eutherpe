package actions

import (
    "internal/vars"
    "internal/dj"
    "net/url"
    "fmt"
)

func RemovePlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-remove request.")
    }
    if dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists) == nil {
        return fmt.Errorf("Playlist '%s' not exists.", playlist[0])
    }
    for p, curr_playlist := range eutherpeVars.Playlists {
        if curr_playlist.Name == playlist[0] {
            eutherpeVars.Playlists = append(eutherpeVars.Playlists[:p], eutherpeVars.Playlists[(p+1):]...)
            if eutherpeVars.RenderedPlaylist == playlist[0] {
                eutherpeVars.RenderedPlaylist = ""
            }
            break
        }
    }
    return nil
}