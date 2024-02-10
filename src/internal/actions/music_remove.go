package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func MusicRemove(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed music-remove request.")
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        songFilePath := GetSongFilePathFromSelectionId(selectionId)
        for n, nextSong := range eutherpeVars.Player.UpNext {
            if nextSong.FilePath == songFilePath {
                eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext[:n], eutherpeVars.Player.UpNext[n+1:]...)
                break
            }
        }
    }
    return nil
}