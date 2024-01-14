package actions

import (
    "internal/vars"
    "internal/dj"
    "net/url"
    "fmt"
)

func AddSelectionToPlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    selection, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed addselectiontoplaylist request.")
    }
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has || len(playlist) != 1 {
        return fmt.Errorf("Malformed addselectiontoplaylist request.")
    }
    editedPlaylist := dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists)
    if editedPlaylist == nil {
        return fmt.Errorf("Playlist %s not found.", playlist[0])
    }
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            return err
        }
        editedPlaylist.Add(song)
    }
    return nil
}
