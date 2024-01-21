package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestMusicClearAll(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, regularJohn, deadMen, stayClean, fever)
    err := MusicClearAll(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicClearAll() has returned an error when it should not.\n")
    } else if len(eutherpeVars.Player.UpNext) != 0 {
        t.Errorf("MusicClearAll() seems not to be clearing the UpNext list accordingly.\n")
    }
    if MusicClearAll(eutherpeVars, userData) != nil {
        t.Errorf("MusicClearAll() with empty UpNext list should be accepted and do not generate any error.\n")
    }
}
