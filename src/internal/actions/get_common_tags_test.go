package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestGetCommonTags(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Sonic Youth"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Sonic Youth"]["Washing Machine"] = []mplayer.SongInfo {
        mplayer.SongInfo { "diamond-sea.mp3", "Diamond Sea", "Sonic Youth", "Washing Machine", "11", "1994", "", "Indie", },
    }
    eutherpeVars.Collection["Nirvana"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Nirvana"]["Unplugged"] = []mplayer.SongInfo {
        mplayer.SongInfo { "about-a-girl.mp3", "About a Girl", "Nirvana", "Unplugged", "01", "1993", "", "Grunge", },
    }
    eutherpeVars.Tags.Add("Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3", "90s", "Indie")
    eutherpeVars.Tags.Add("Nirvana/Unplugged/About a Girl:about-a-girl.mp3", "90s", "Grunge")
    userData := &url.Values{}
    err := GetCommonTags(eutherpeVars, userData)
    if err == nil {
        t.Errorf("GetCommonTags() did not return an error when it should.\n")
    } else if err.Error() != "Malformed get-commontags request." {
        t.Errorf("GetCommonTags() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3\", \"Nirvana/Unplugged/About a Girl:about-a-girl.mp3\" ]")
    err = GetCommonTags(eutherpeVars, userData)
    if err != nil {
        t.Errorf("GetCommonTags() has returned an error when it must not : '%s'\n", err.Error())
    }
    if len(eutherpeVars.LastCommonTags) != 1 {
        t.Errorf("GetCommonTags() has returned more than one common tag.\n")
    } else if eutherpeVars.LastCommonTags[0] != "90s" {
        t.Errorf("GetCommonTags() has returned unexpected common tag : '%s'\n", eutherpeVars.LastCommonTags)
    }
    if eutherpeVars.CurrentConfig != "RemoveTags" {
        t.Errorf("GetCommonTags() did not set eutherpeVars.CurrentConfig accordingly.\n")
    }
    if eutherpeVars.LastSelection != "diamond-sea.mp3,about-a-girl.mp3" {
        t.Errorf("GetCommonTags() did not set eutherpeVars.LastSelection accordingly : '%s'\n", eutherpeVars.LastSelection)
    }
}
