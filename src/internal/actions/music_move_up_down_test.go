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

func TestMusicMoveUp(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, regularJohn, deadMen, stayClean, fever)
    err := MusicMoveUp(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicMoveUp() has succeeded when it should fail.\n")
    } else if err.Error() != "Malformed music-moveup request." {
        t.Errorf("MusicMoveUp() has returned an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Motorhead/Overkill/Stay Clean:stay-clean.mp3\", \"The Cramps/Songs The Lord Taught Us/Fever:fever.mp3\" ]")
    err = MusicMoveUp(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicMoveUp() has returned an error when it should not.\n")
    } else if len(eutherpeVars.Player.UpNext) != 4 {
        t.Errorf("MusicMoveUp() seems to be corrupting the UpNext selection.\n")
    } else if eutherpeVars.Player.UpNext[0] != regularJohn ||
              eutherpeVars.Player.UpNext[1] != stayClean ||
              eutherpeVars.Player.UpNext[2] != fever ||
              eutherpeVars.Player.UpNext[3] != deadMen {
        t.Errorf("MusicMoveUp() dynamics seems broken.\n")
    }
}

func TestMusicMoveDown(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, regularJohn, deadMen, stayClean, fever)
    err := MusicMoveDown(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicMoveDown() has succeeded when it should fail.\n")
    } else if err.Error() != "Malformed music-movedown request." {
        t.Errorf("MusicMoveDown() has returned an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Queens Of The Stone Age/Queens Of The Stone Age/Regular John:regular-john.mp3\", \"Motorhead/Stay Clean/Stay Clean:stay-clean.mp3\" ]")
    err = MusicMoveDown(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicMoveDown() has returned an error when it should not.\n")
    } else if len(eutherpeVars.Player.UpNext) != 4 {
        t.Errorf("MusicMoveDown() seems to be corrupting the UpNext selection.\n")
    } else if eutherpeVars.Player.UpNext[0] != deadMen ||
              eutherpeVars.Player.UpNext[1] != regularJohn ||
              eutherpeVars.Player.UpNext[2] != fever ||
              eutherpeVars.Player.UpNext[3] != stayClean {
        t.Errorf("MusicMoveDown() dynamics seems broken.\n")
    }
}
