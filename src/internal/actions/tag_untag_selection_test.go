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

func TestTagSelection(t *testing.T) {
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
    userData := &url.Values{}
    err := TagSelection(eutherpeVars, userData)
    if err == nil {
        t.Errorf("TagSelection() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-tagselectionas/collection-untagselections request." {
        t.Errorf("TagSelection() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[\"Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3\",\"Nirvana/Unplugged/About a Girl:about-a-girl.mp3\"]")
    err = TagSelection(eutherpeVars, userData)
    if err == nil {
        t.Errorf("TagSelection() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-tagselectionas/collection-untagselections request." {
        t.Errorf("TagSelection() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldTags, "90s")
    err = TagSelection(eutherpeVars, userData)
    if err != nil {
        t.Errorf("TagSelection() did return an error when it should not.\n")
    } else {
        tags := eutherpeVars.Tags.Get("90s")
        if len(tags) != 2 {
            t.Errorf("TagSelection() seems not be working as it should.\n")
        } else if tags[0] != "Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3" ||
                  tags[1] != "Nirvana/Unplugged/About a Girl:about-a-girl.mp3" {
            t.Errorf("TagSelection() is not tagging accordingly.\n")
        }
    }
}

func TestUntagSelection(t *testing.T) {
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
    if len(eutherpeVars.Tags.Get("90s")) != 2 {
        t.Errorf("Songs were not pre-tagged accordingly.\n")
    }
    userData := &url.Values{}
    err := UntagSelection(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UntagSelection() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-tagselectionas/collection-untagselections request." {
        t.Errorf("UntagSelection() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[\"Sonic Youth/Washing Machine/Diamond Sea:diamond-sea.mp3\",\"Nirvana/Unplugged/About a Girl:about-a-girl.mp3\"]")
    err = UntagSelection(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UntagSelection() did not return an error when it should.\n")
    } else if err.Error() != "Malformed collection-tagselectionas/collection-untagselections request." {
        t.Errorf("UntagSelection() did return an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldTags, "90s")
    err = UntagSelection(eutherpeVars, userData)
    if err != nil {
        t.Errorf("UntagSelection() did return an error when it should not.\n")
    } else {
        tags := eutherpeVars.Tags.Get("90s")
        if len(tags) != 0 {
            t.Errorf("UntagSelection() seems not be working as it should.\n")
        }
    }
}
