package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
    "flag"
    "strconv"
)

func MusicSetVolume(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../mplayer"
    }
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
    mplayer.SetVolume(value, customPath)
    eutherpeVars.Player.VolumeLevel = uint(value)
    return nil
}
