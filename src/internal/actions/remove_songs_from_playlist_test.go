package actions

import (
    "internal/vars"
    "internal/mplayer"
    "internal/dj"
    "net/url"
    "testing"
)

func TestRemoveSongsFromPlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Collection = make(mplayer.MusicCollection)
    eutherpeVars.Collection["Queens Of The Stone Age"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Motorhead"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["The Cramps"] = make(map[string][]mplayer.SongInfo)
    eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", },
    }
    eutherpeVars.Collection["Motorhead"]["Bomber"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["Motorhead"]["Overkill"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", },
    }
    eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"] = []mplayer.SongInfo {
        mplayer.SongInfo { "/fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", },
    }
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "speed-metal-e-psychobilly-do-bom" })
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist { Name: "stoner-rock" })
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Overkill"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Motorhead"]["Bomber"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["The Cramps"]["Songs The Lord Taught Us"][0])
    eutherpeVars.Playlists[0].Add(eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    eutherpeVars.Playlists[1].Add(eutherpeVars.Collection["Queens Of The Stone Age"]["Queens Of The Stone Age"][0])
    userData := &url.Values{}
    err := RemoveSongsFromPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("RemoveSongsFromPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-removesongs request." {
        t.Errorf("RemoveSongsFromPlaylist() did return an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldSelection, "[ \"boooooo!\" ]")
    err = RemoveSongsFromPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("RemoveSongsFromPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-removesongs parameter." {
        t.Errorf("RemoveSongsFromPlaylist() did return an unexpected error.\n")
    }
    userData.Del(vars.EutherpePostFieldSelection)
    userData.Add(vars.EutherpePostFieldSelection, "[ \"speed-metal-e-psychobilly-do-bom:QUeens Of The Stone Age/Queens Of The Stone Age/Regular John:/regular-john.mp3\", \"speed-metal-e-psychobilly-do-bom:Motorhead/Bomber:/dead_men_tell_no_tales.mp3\" ]")
    err = RemoveSongsFromPlaylist(eutherpeVars, userData)
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("/stay-clean.mp3") != 0 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("/dead_men_tell_no_tales.mp3") != -1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("/fever.mp3") != 1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("/regular-john.mp3") != -1 ||
       eutherpeVars.Playlists[1].GetSongIndexByFilePath("/regular-john.mp3") != 0 {
        t.Errorf("RemoveSongsFromPlaylist() seems not to be manipulating the playlist as expected : '%v', '%v'\n", eutherpeVars.Playlists[0],
                                                                                                                   eutherpeVars.Playlists[1])
    }
}
