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
    "testing"
    "os"
)

func TestBackupRestorePlaylists(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    os.MkdirAll("/tmp/mdev", 0777)
    os.MkdirAll("/tmp/playlists", 0777)
    defer os.RemoveAll("/tmp/mdev")
    defer os.RemoveAll("/tmp/playlists")
    err := BackupPlaylists(eutherpeVars, nil)
    if err == nil {
        t.Errorf("BackupPlaylists() is not failing when it should.")
    } else if err.Error() != "There is no playlists to backup." {
        t.Errorf("BackupPlaylists() has returned unexpected error.")
    }
    eutherpeVars.ConfHome = "/tmp"
    littleDoll := mplayer.SongInfo { "little-doll.mp3", "Little Doll", "The Stooges", "The Stooges", "4", "1969", "", "Proto-punk", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    stoogesPlaylist := dj.Playlist {}
    stoogesPlaylist.Name = "Stooges"
    stoogesPlaylist.Add(littleDoll)
    rockPlaylist := dj.Playlist{}
    rockPlaylist.Name = "Rock"
    rockPlaylist.Add(littleDoll)
    rockPlaylist.Add(deadMen)
    rockPlaylist.Add(regularJohn)
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, stoogesPlaylist, rockPlaylist)
    err = BackupPlaylists(eutherpeVars, nil)
    if err == nil {
        t.Errorf("BackupPlaylists() is not failing when it should.")
    } else if err.Error() != "No storage device was set." {
        t.Errorf("BackupPlaylists() has returned unexpected error.")
    }
    eutherpeVars.CachedDevices.MusicDevId = "/tmp/mdev"
    err = BackupPlaylists(eutherpeVars, nil)
    if err != nil {
        t.Errorf("BackupPlaylists() is failing when it should not.\n")
    }
    if _, err := os.Stat("/tmp/mdev/.eutherpe/playlists/Stooges"); err != nil {
        t.Errorf("Stooges playlist was not copied.")
    }
    if _, err := os.Stat("/tmp/mdev/.eutherpe/playlists/Rock"); err != nil {
        t.Errorf("Rock playlist was not copied.")
    }
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    eutherpeVars.CachedDevices.MusicDevId = ""
    err = RestorePlaylists(eutherpeVars, nil)
    if err == nil {
        t.Errorf("RestorePlaylists() is not failing when it should.")
    } else if err.Error() != "No storage device was set." {
        t.Errorf("RestorePlaylists() has returned unexpected error.")
    }
    eutherpeVars.CachedDevices.MusicDevId = "/tmp/mdev"
    err = RestorePlaylists(eutherpeVars, nil)
    if err != nil {
        t.Errorf("RestorePlaylists() is failing when it should not.\n")
    }
    if len(eutherpeVars.Playlists) != 2 {
        t.Errorf("Playlists slice has unexpected size after restoring.\n")
    }
    if eutherpeVars.Playlists[0].Name != "Rock" {
        t.Errorf("Rock playlists seems not restored accordingly.")
    }
    if eutherpeVars.Playlists[1].Name != "Stooges" {
        t.Errorf("Stooges playlists seems not restored accordingly.")
    }
}
