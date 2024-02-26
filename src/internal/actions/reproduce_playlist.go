package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func ReproducePlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-reproduce request.")
    }
    err := fmt.Errorf("Playlist '%s' not found!", playlist[0])
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    for _, currPlaylist := range eutherpeVars.Playlists {
        if currPlaylist.Name == playlist[0] {
            if len(eutherpeVars.Player.UpNext) > 0 {
                eutherpeVars.Unlock()
                MusicClearAll(eutherpeVars, nil)
                eutherpeVars.Lock()
            }
            eutherpeVars.Player.UpNext = currPlaylist.Songs()
            eutherpeVars.Unlock()
            err = MusicPlay(eutherpeVars, nil)
            eutherpeVars.Lock()
            break
        }
    }
    return err
}
