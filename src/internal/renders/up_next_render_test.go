//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "internal/mplayer"
    "fmt"
    "testing"
)

func TestUpNextRender(t *testing.T) {
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
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["Motorhead"]["Bomber"][0])
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleUpNext)
    output := UpNextRender(templatedInput, eutherpeVars)
    if output != "<ul class=\"nested\"><input type=\"checkbox\" id=\"Motorhead/Bomber:dead_men_tell_no_tales.mp3\" class=\"UpNext\">Dead Men Tell No Tales<br><input type=\"checkbox\" id=\"The Cramps/Songs The Lord Taught Us:fever.mp3\" class=\"UpNext\">Fever<br><input type=\"checkbox\" id=\"Queens Of The Stone Age/Queens Of The Stone Age:regular-john.mp3\" class=\"UpNext\">Regular John<br></ul>" {
        t.Errorf("UpNextRender() seems not to be working accordingly.\n")
    }
}
