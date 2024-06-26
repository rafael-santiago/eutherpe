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

func TestAddSelectionToNext(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Queens Of The Stone Age"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Motorhead"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["The Cramps"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []mplayer.SongInfo {
        mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    eutherpeVars.Collection["Motorhead"]["Bomber"] = []mplayer.SongInfo {
        mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["Motorhead"]["Overkill"] = []mplayer.SongInfo {
        mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"] = []mplayer.SongInfo {
        mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    userData := &url.Values{}
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Motorhead/Overkill/Stay Clean:stay-clean.mp3\", \"The Cramps/Songs The Lord Taught Us/Fever:fever.mp3\", \"Queens Of The Stone Age/Queens Of The Stone Age/Regular John:regular-john.mp3\" ]")
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["Motorhead"]["Bomber"][0])
    err := AddSelectionToNext(eutherpeVars, userData)
    if err != nil {
        t.Errorf("AddSelectionToNext() returned error.\n")
    }
    if len(eutherpeVars.Player.UpNext) != 4 {
        t.Errorf("eutherpeVars.Player.UpNext has wrong total of songs.\n")
    }
    if eutherpeVars.Player.UpNext[0] != eutherpeVars.Collection["Motorhead"]["Bomber"][0] {
        t.Errorf("eutherpeVars.Player.UpNext seems like not following the order of userData.\n")
    }
    if eutherpeVars.Player.UpNext[1] != eutherpeVars.Collection["Motorhead"]["Overkill"][0] {
        t.Errorf("eutherpeVars.Player.UpNext seems like not following the order of userData.\n")
    }
    if eutherpeVars.Player.UpNext[2] != eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0] {
        t.Errorf("eutherpeVars.Player.UpNext seems like not following the order of userData.\n")
    }
    if eutherpeVars.Player.UpNext[3] != eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0] {
        t.Errorf("eutherpeVars.Player.UpNext seems like not following the order of userData.\n")
    }
    userData = &url.Values{}
    err = AddSelectionToNext(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToNext did not return an error.\n")
    }
    if err.Error() != "Malformed addselectiontonext request." {
        t.Errorf("Unexpected error message.\n")
    }
    userData = &url.Values{}
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Motorhead/Overkill/Stay Clean:stay-clean.mp3\", \"The Grumpies/Songs The Lord Taught Us/Fever:fever.mp3\", \"Queens Of The Stone Age/Queens Of The Stone Age/Regular John:regular-john.mp3\" ]")
    err = AddSelectionToNext(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToNext did not return an error.\n")
    }
    if err.Error() != "No collection for The Grumpies." {
        t.Errorf("Unexpected error message.\n")
    }
}
