package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func AddSelectionToNext(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed addselectiontonext request.")
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            fmt.Println(err)
            return err
        }
        eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, song)
    }
    return nil
}
