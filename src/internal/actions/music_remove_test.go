package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestMusicRemove(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, regularJohn, deadMen, stayClean, fever)
    err := MusicRemove(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicRemove() has not returned an error when it should.\n")
    } else if err.Error() != "Malformed music-remove request." {
        t.Errorf("MusicRemove() has returned an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldSelection, "Motorhead/Bomber/Dead Men Tell No Tales:dead_men_tell_no_tales.mp3")
    userData.Add(vars.EutherpePostFieldSelection, "Queens Of The Stone Age/Queens Of The Stone Age/Regular John:regular-john.mp3")
    userData.Add(vars.EutherpePostFieldSelection, "The Cramps/Songs The Lord Taught Us/Fever:fever.mp3")
    err = MusicRemove(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicRemove() has returned an error when it should not.\n")
    } else if len(eutherpeVars.Player.UpNext) != 1 {
        t.Errorf("MusicRemove() seems not to be removing the selection list well.\n")
    } else if eutherpeVars.Player.UpNext[0] != stayClean {
        t.Errorf("MusicRemove() seems to be removing the unselected song.\n")
    }
}
