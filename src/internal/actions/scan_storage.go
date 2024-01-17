package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
)

func ScanStorage(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return fmt.Errorf("Unset MusicDevId.")
    }
    newCollection, err := mplayer.LoadMusicCollection(eutherpeVars.CachedDevices.MusicDevId)
    if err != nil {
        return err
    }
    eutherpeVars.Collection = newCollection
    return nil
}
