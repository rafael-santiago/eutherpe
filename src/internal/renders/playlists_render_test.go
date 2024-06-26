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
    "internal/dj"
    "fmt"
    "testing"
)

func TestPlaylistsRender(t *testing.T) {
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
        mplayer.SongInfo { "overkill.mp3", "Overkill", "Motorhead", "Overkill", "1", "1979", "", "Speed Metal", },
        mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
        mplayer.SongInfo { "pay-your-price.mp3", "(I Won't) Pay Your Price", "Motorhead", "Overkill", "3", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"] = []mplayer.SongInfo {
        mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    stoner := dj.Playlist { Name: "Stoner Rock" }
    stoner.Add(eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    speedMetalAndPsychobilly := dj.Playlist { Name: "Speed Metal & Psychobilly" }
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][2])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][0])
    speedMetalAndPsychobilly.Add(eutherpeVars.Collection["Motorhead"]["Overkill"][1])
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, stoner, speedMetalAndPsychobilly)
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedlePlaylists)
    output := PlaylistsRender(templatedInput, eutherpeVars)
    if output != "<ul id=\"eutherpeUL\"><li><input type=\"checkbox\" onclick=\"flush_child(this);selectPlaylist(this);\" id=\"Stoner Rock\" class=\"Playlist\"><span class=\"caret\">Stoner Rock</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"Stoner Rock:Queens Of The Stone Age/Queens Of The Stone Age/:regular-john.mp3\" class=\"PlaylistSong\">Regular John</li></ul></li></ul><ul id=\"eutherpeUL\"><li><input type=\"checkbox\" onclick=\"flush_child(this);selectPlaylist(this);\" id=\"Speed Metal & Psychobilly\" class=\"Playlist\"><span class=\"caret\">Speed Metal & Psychobilly</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"Speed Metal & Psychobilly:The Cramps/Songs The Lord Taught Us/:fever.mp3\" class=\"PlaylistSong\">Fever</li><li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"Speed Metal & Psychobilly:Motorhead/Overkill/:pay-your-price.mp3\" class=\"PlaylistSong\">(I Won't) Pay Your Price</li><li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"Speed Metal & Psychobilly:Motorhead/Overkill/:overkill.mp3\" class=\"PlaylistSong\">Overkill</li><li><input type=\"checkbox\" onclick=\"flush_child(this);setUncheckedAllSongsOutFromPlaylist(this);\" id=\"Speed Metal & Psychobilly:Motorhead/Overkill/:stay-clean.mp3\" class=\"PlaylistSong\">Stay Clean</li></ul></li></ul>" {
        t.Errorf("PlaylistRender() seems not to be rendering accordingly.\n")
    }
}
