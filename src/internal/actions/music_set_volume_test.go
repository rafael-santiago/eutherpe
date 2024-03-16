package actions

import (
    "internal/vars"
    "testing"
    "net/url"
)

func TestMusicSetVolume(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := MusicSetVolume(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicSetVolume() did not return an error when it should.\n")
    } else if err.Error() != "Malformed music-setvolume request." {
        t.Errorf("MusicSetVolume() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldVolumeLevel, "XXI")
    err = MusicSetVolume(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicSetVolume() did not return an error when it should.\n")
    } else if err.Error() != "strconv.Atoi: parsing \"XXI\": invalid syntax" {
        t.Errorf("MusicSetVolume() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldVolumeLevel)
    userData.Add(vars.EutherpePostFieldVolumeLevel, "21")
    err = MusicSetVolume(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicSetVolume() did return an error when it should not.\n")
    } else if eutherpeVars.Player.VolumeLevel != 21 {
        t.Errorf("MusicSetVolume() is not saving set volume level accordingly.\n")
    }
}
