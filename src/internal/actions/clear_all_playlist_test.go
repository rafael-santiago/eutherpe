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
    "internal/dj"
    "internal/mplayer"
    "net/url"
    "testing"
)

func TestClearAllPlaylist(t *testing.T) {
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
    err := ClearAllPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ClearAllPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Malformed playlist-clearall request." {
        t.Errorf("ClearAllPlaylist() did return an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "Papai Noel")
    err = ClearAllPlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("ClearAllPlaylist() did not return an error when it should.\n")
    } else if err.Error() != "Playlist 'Papai Noel' not exists." {
        t.Errorf("ClearAllPlaylist() did return an unexpected error.\n")
    }
    userData.Del(vars.EutherpePostFieldPlaylist)
    userData.Add(vars.EutherpePostFieldPlaylist, "speed-metal-e-psychobilly-do-bom")
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("fever.mp3") == -1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") == -1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("dead_men_tell_no_tales.mp3") == -1 {
        t.Errorf("Test playlist seems corrupted : '%v'\n", eutherpeVars.Playlists[0])
    }
    err = ClearAllPlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ClearAllPlaylist() did return an error when it should not.\n")
    }
    if eutherpeVars.Playlists[0].GetSongIndexByFilePath("fever.mp3") != -1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("stay-clean.mp3") != -1 ||
       eutherpeVars.Playlists[0].GetSongIndexByFilePath("dead_men_tell_no_tales.mp3") != -1 {
        t.Errorf("ClearAllPlaylist() seems not to be working accordingly : '%v'\n", eutherpeVars.Playlists[0])
    }
}
