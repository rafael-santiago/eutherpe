package actions

import (
    "internal/vars"
    "internal/dj"
    "net/url"
    "fmt"
)

func ClearAllPlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-clearall request.")
    }
    editedPlaylist := dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists)
    if editedPlaylist == nil {
        return fmt.Errorf("Playlist '%s' not exists.", playlist[0])
    }
    editedPlaylist.ClearAll()
    return nil
}
