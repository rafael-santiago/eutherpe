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

func TestCollectionRender(t *testing.T) {
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
    templatedInput := fmt.Sprintf("%s", vars.EutherpeTemplateNeedleCollection)
    output := CollectionRender(templatedInput, eutherpeVars)
    if output != "<ul id=\"eutherpeUL\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead\" class=\"CollectionArtist\"><span class=\"caret\">Motorhead</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Overkill-eutpid_0\" class=\"CollectionAlbum\"><span class=\"caret\">Overkill</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Overkill-eutpid_0/Overkill-eutpid_1:overkill.mp3\" class=\"CollectionSong\">Overkill</li><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Overkill-eutpid_0/Stay Clean-eutpid_2:stay-clean.mp3\" class=\"CollectionSong\">Stay Clean</li><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Overkill-eutpid_0/(I Won't) Pay Your Price-eutpid_3:pay-your-price.mp3\" class=\"CollectionSong\">(I Won't) Pay Your Price</li></ul></li><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Bomber-eutpid_4\" class=\"CollectionAlbum\"><span class=\"caret\">Bomber</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Motorhead/Bomber-eutpid_4/Dead Men Tell No Tales-eutpid_5:dead_men_tell_no_tales.mp3\" class=\"CollectionSong\">Dead Men Tell No Tales</li></ul></li></ul></li><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Queens Of The Stone Age\" class=\"CollectionArtist\"><span class=\"caret\">Queens Of The Stone Age</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Queens Of The Stone Age/Queens Of The Stone Age-eutpid_6\" class=\"CollectionAlbum\"><span class=\"caret\">Queens Of The Stone Age</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"Queens Of The Stone Age/Queens Of The Stone Age-eutpid_6/Regular John-eutpid_7:regular-john.mp3\" class=\"CollectionSong\">Regular John</li></ul></li></ul></li><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"The Cramps\" class=\"CollectionArtist\"><span class=\"caret\">The Cramps</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"The Cramps/Songs The Lord Taught Us-eutpid_8\" class=\"CollectionAlbum\"><span class=\"caret\">Songs The Lord Taught Us</span><ul class=\"nested\"><li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"The Cramps/Songs The Lord Taught Us-eutpid_8/Fever-eutpid_9:fever.mp3\" class=\"CollectionSong\">Fever</li></ul></li></ul></li></ul>" {
        t.Errorf("CollectionRender() seems not to be working accordingly. : '%s'\n", output)
    }
}
