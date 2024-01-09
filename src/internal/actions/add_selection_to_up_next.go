package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
)

func AddSelectionToUpNext(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    selection, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed addselectiontoupnext request.")
    }
    var upNextNewHead []mplayer.SongInfo
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            return err
        }
        upNextNewHead = append(upNextNewHead, song)
    }
    eutherpeVars.Player.UpNext = append(upNextNewHead, eutherpeVars.Player.UpNext...)
    return nil
}
