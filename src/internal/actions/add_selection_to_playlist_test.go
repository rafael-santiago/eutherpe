package actions

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "net/url"
    "testing"
)

func TestAddSelectionToPlaylist(t *testing.T) {
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
    err := AddSelectionToPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToPlaylist() should return an error.\n")
    }
    if err.Error() != "Malformed addselectiontoplaylist request." {
        t.Errorf("AddSelectionToPlaylist() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"Motorhead/Overkill/Stay Clean:stay-clean.mp3\", \"Queens Of The Stone Age/Queens Of The Stone Age/Regular John:regular-john.mp3\" ]")
    err = AddSelectionToPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToPlaylist() should return an error.\n")
    }
    if err.Error() != "Malformed addselectiontoplaylist request." {
        t.Errorf("AddSelectionToPlaylist() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "1-2-3-eh-um-teste")
    err = AddSelectionToPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToPlaylist() should return an error.\n")
    }
    if err.Error() != "Playlist 1-2-3-eh-um-teste not found." {
        t.Errorf("AddSelectionToPlaylist() has returned an unexpected error : '%s'.\n", err.Error())
    }
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist{ Name: "1-2-3-eh-um-teste" })
    err = AddSelectionToPlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("AddSelectionToPlaylist() has failed while it should not.\n")
    }
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") != 0 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("regular-john.mp3") != 1 {
        t.Errorf("AddSelectionToPlaylist() seems broken because it has not follow the selection order.\n")
    }
    userData.Del(vars.EutherpePostFieldSelection)
    userData.Add(vars.EutherpePostFieldSelection, "[ \"The Cramps/Songs The Lord Taught Us/Fever:fever.mp3\" ]")
    err = AddSelectionToPlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("AddSelectionToPlaylist() has failed while it should not.\n")
    }
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") != 0 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("regular-john.mp3") != 1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("fever.mp3") != 2 {
        t.Errorf("AddSelectionToPlaylist() seems broken because it has not follow the selection order.\n")
    }
    userData.Del(vars.EutherpePostFieldSelection)
    userData.Add(vars.EutherpePostFieldSelection, "[ \"The Cramps/Songs The Lord Taught Us/Garbage Man:garbage-man.mp3\" ]")
    err = AddSelectionToPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("AddSelectionToPlaylist() has worked when it should not.\n")
    }
    if err.Error() != "No song garbage-man.mp3 in album Songs The Lord Taught Us by The Cramps was found." {
        t.Error("AddSelectionToPlaylist() has returned an unexpected error.\n")
    }
}


