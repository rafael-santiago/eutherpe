package actions

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "net/url"
    "testing"
)

func TestMoveUpPlaylistSongs(t *testing.T) {
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
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "speed-metal-e-psychobilly-do-bom" })
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Overkill"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Bomber"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    userData := &url.Values{}
    err := MoveUpPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveUpPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Malformed playlist-moveup/down request." {
        t.Errorf("MoveUpPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "BlauQueNaoExisteTavaDoidao...")
    err = MoveUpPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveUpPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Malformed playlist-moveup/down request." {
        t.Errorf("MoveUpPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"speed-metal-e-psychobilly-do-bom:Motorhead/Bomber/Dead Men Tell No Tales:dead_men_tell_no_tales.mp3\", \"speed-metal-e-psychobilly-do-bom:The Grumpies/Songs The Lord Taught Us/Fever:fever.mp3\" ]")
    err = MoveUpPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveUpPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Playlist 'BlauQueNaoExisteTavaDoidao...' not exists." {
        t.Errorf("MoveUpPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "speed-metal-e-psychobilly-do-bom")
    err = MoveUpPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveUpPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "No collection for The Grumpies." {
        t.Errorf("MoveUpPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldSelection)
    userData.Add(vars.EutherpePostFieldSelection, "[ \"speed-metal-e-psychobilly-do-bom:Motorhead/Bomber/Dead Men Tell No Tales:dead_men_tell_no_tales.mp3\", \"speed-metal-e-psychobilly-do-bom:The Cramps/Songs The Lord Taught Us/Fever:fever.mp3\" ]")
    err = MoveUpPlaylistSongs(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MoveUpPlaylistSongs() has return an error when it should not.\n")
    }
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("dead_men_tell_no_tales.mp3") != 0 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("fever.mp3") != 1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") != 2 {
        t.Errorf("MoveUpPlaylistSongs() seems not being manipulating playlist accordingly : '%v'\n", eutherpeVars.Playlists[0])
    }
}

func TestMoveDownPlaylistSongs(t *testing.T) {
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
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "speed-metal-e-psychobilly-do-bom" })
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Overkill"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Bomber"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    userData := &url.Values{}
    err := MoveDownPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveDownPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Malformed playlist-moveup/down request." {
        t.Errorf("MoveDownPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "BlauQueNaoExisteTavaDoidao...")
    err = MoveDownPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveDownPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Malformed playlist-moveup/down request." {
        t.Errorf("MoveDownPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"speed-metal-e-psychobilly-do-bom:Motorhead/Bomber/Dead Men Tell No Tales:dead_men_tell_no_tales.mp3\", \"speed-metal-e-psychobilly-do-bom:Moforhead/Overkill/Stay Clean:stay-clean.mp3\" ]")
    err = MoveDownPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveDownPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "Playlist 'BlauQueNaoExisteTavaDoidao...' not exists." {
        t.Errorf("MoveDownPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "speed-metal-e-psychobilly-do-bom")
    err = MoveDownPlaylistSongs(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MoveDownPlaylistSongs() has not return an error while it should.\n")
    } else if err.Error() != "No collection for Moforhead." {
        t.Errorf("MoveDownPlaylistSongs() has returned an unexpected error : '%s'\n", err.Error())
    }
    userData.Del(vars.EutherpePostFieldSelection)
    userData.Add(vars.EutherpePostFieldSelection, "[ \"speed-metal-e-psychobilly-do-bom:Motorhead/Bomber/Dead Men Tell No Tales:dead_men_tell_no_tales.mp3\", \"speed-metal-e-psychobilly-do-bom:Motorhead/Overkill/Stay Clean:stay-clean.mp3\" ]")
    err = MoveDownPlaylistSongs(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MoveDownPlaylistSongs() has return an error when it should not.\n")
    }
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("fever.mp3") != 2 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") != 0 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("dead_men_tell_no_tales.mp3") != 1 {
        t.Errorf("MoveDownPlaylistSongs() seems not being manipulating playlist accordingly : '%v'\n", eutherpeVars.Playlists[0])
    }
}
