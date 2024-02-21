package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
    "strconv"
)

func MusicSetVolume(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    volumeLevel, has := (*userData)[vars.EutherpePostFieldVolumeLevel]
    if !has {
        return fmt.Errorf("Malformed music-setvolume request.")
    }
    value, err := strconv.Atoi(volumeLevel[0])
    if err != nil {
        return err
    }
    mplayer.SetVolume(value)
    eutherpeVars.Player.VolumeLevel = uint(value)
    return nil
}
