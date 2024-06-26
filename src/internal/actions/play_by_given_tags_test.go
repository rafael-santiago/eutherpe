//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestPlayByGivenTags(t *testing.T) {
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
    err := PlayByGivenTags(eutherpeVars, userData)
    if err == nil {
        t.Errorf("PlayByGivenTags() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-playbygiventags request." {
        t.Errorf("PlayByGivenTags() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldTags, "GRUNGE")
    err = PlayByGivenTags(eutherpeVars, userData)
    if err == nil {
        t.Errorf("PlayByGivenTags() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-playbygiventags request." {
        t.Errorf("PlayByGivenTags() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldAmount, "2")
    err = PlayByGivenTags(eutherpeVars, userData)
    if err != nil {
        t.Errorf("PlayByGivenTags() has returned an error when it must not : '%s'\n", err.Error())
    }
    if len(eutherpeVars.Player.UpNext) != 1 {
        t.Errorf("PlayByGivenTags() is not playing as it should.\n")
    }
    if eutherpeVars.Player.NowPlaying.FilePath != "about-a-girl.mp3" {
        t.Errorf("PlayByGivenTags() is not playing the expected song.\n")
    }
    MusicStop(eutherpeVars, nil)
    userData.Del(vars.EutherpePostFieldTags)
    userData.Add(vars.EutherpePostFieldTags, " 90s , Grunge ")
    err = PlayByGivenTags(eutherpeVars, userData)
    if err != nil {
        t.Errorf("PlayByGivenTags() has returned an error when it must not : '%s'\n", err.Error())
    }
    if len(eutherpeVars.Player.UpNext) != 2 {
        t.Errorf("PlayByGivenTags() is not playing as it should.\n")
    }
    if (eutherpeVars.Player.UpNext[0].FilePath != "about-a-girl.mp3" &&
        eutherpeVars.Player.UpNext[1].FilePath != "about-a-girl.mp3") ||
       (eutherpeVars.Player.UpNext[0].FilePath != "diamond-sea.mp3" &&
        eutherpeVars.Player.UpNext[1].FilePath != "diamond-sea.mp3") {
        t.Errorf("PlayByGivenTags() did not select the expected song.\n")
    }
    MusicStop(eutherpeVars, nil)
}
