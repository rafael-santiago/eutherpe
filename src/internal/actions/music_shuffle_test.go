package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestMusicShuffle(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, regularJohn, deadMen, stayClean, fever)
    originalUpNext := eutherpeVars.Player.UpNext
    err := MusicShuffle(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicShuffle() has returned an error when it should not.\n")
    } else if len(eutherpeVars.Player.UpNext) != 4 {
        t.Errorf("MusicShuffle() seems to be removing songs from UpNext.\n")
    } else if eutherpeVars.Player.UpNext[0] == regularJohn &&
              eutherpeVars.Player.UpNext[1] == deadMen &&
              eutherpeVars.Player.UpNext[2] == stayClean &&
              eutherpeVars.Player.UpNext[3] == fever {
        t.Errorf("MusicShuffle() seems not to be changing the UpNext ordering.\n")
    } else if !found(regularJohn, eutherpeVars.Player.UpNext) ||
              !found(deadMen, eutherpeVars.Player.UpNext) ||
              !found(stayClean, eutherpeVars.Player.UpNext) ||
              !found(fever, eutherpeVars.Player.UpNext) {
        t.Errorf("MusicShuffle() seems not to be selecting all songs from the original selection.\n")
    } else if !eutherpeVars.Player.Shuffle {
        t.Errorf("MusicShuffle() seems not to be setting the Shuffle flag accordingly.\n")
    } else if !isEqual(eutherpeVars.Player.UpNextBkp, originalUpNext) {
        t.Errorf("MusicShuffle() seems not to be saving the original UpNext selection.\n")
    } else {
        err = MusicShuffle(eutherpeVars, userData)
        if err != nil {
            t.Errorf("MusicShuffle() has returned an error when it should not.\n")
        } else if !testIsEqual(eutherpeVars.Player.UpNext, originalUpNext) {
            t.Errorf("MusicShuffle() seems not to be restoring the original UpNext selection.\n")
        } else if eutherpeVars.Player.Shuffle {
            t.Errorf("MusicShuffle() seems not to be unsetting the Shuffle flag accordingly.\n")
        } else if len(eutherpeVars.Player.UpNextBkp) != 0 {
            t.Errorf("MusicShuffle() seems not to be zeroing the UpNextBkp accordingly.\n")
        }
    }
}

func found(song mplayer.SongInfo, songs []mplayer.SongInfo) bool {
    for _, curr_song := range songs {
        if curr_song == song {
            return true
        }
    }
    return false
}

func testIsEqual(a, b[]mplayer.SongInfo) bool {
    if len(a) != len(b) {
        return false
    }
    for x, _ := range a {
        if a[x] != b[x] {
            return false
        }
    }
    return true
}
