package actions

import (
    "internal/vars"
    "internal/dj"
    "internal/mplayer"
    "net/url"
    "fmt"
    "strings"
)

func RemoveSongsFromPlaylist(eutherpeVars *vars.EutherpeVars,
                             userData *url.Values) error {
    selections, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed playlist-removesongs request.")
    }
    for _, selection := range selections {
        data := strings.Split(selection, ":")
        if len(data) != 2 {
            return fmt.Errorf("Malformed playlist-removesongs parameter.")
        }
        var editedPlaylist *dj.Playlist
        editedPlaylist = nil
        for p, _ := range eutherpeVars.Playlists {
            if eutherpeVars.Playlists[p].Name == data[0] {
                editedPlaylist = &eutherpeVars.Playlists[p]
                break
            }
        }
        if editedPlaylist == nil {
            return fmt.Errorf("Playlists '%s' does not exist.", data[0])
        }
        editedPlaylist.Remove(mplayer.SongInfo{ FilePath: data[1] })
    }
    return nil
}
