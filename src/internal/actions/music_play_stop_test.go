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
    "internal/mplayer"
    "net/url"
    "time"
    "testing"
    "fmt"
)

func TestMusicPlayStop(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := MusicPlay(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicPlay() has not returned an error when it should.\n")
    } else if err.Error() != "There is no selection to play." {
        t.Errorf("MusicPlay() has returned an unexpected error : '%s'.\n", err.Error())
    }
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, fever, stayClean, deadMen, regularJohn)
    err = MusicPlay(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicPlay() has returned an error when it should not.\n")
    } else if eutherpeVars.Player.Stopped {
        t.Errorf("MusicPlay() seems not to be setting Player.Stopped flag accordingly.\n")
    } else if eutherpeVars.Player.UpNextCurrentOffset > 0 {
        t.Errorf("MusicPlay() seems not to be managing Player.UpNextCurrentOffset counter accordingly.\n")
    } else if eutherpeVars.Player.NowPlaying != fever {
        t.Errorf("MusicPlay() seems to be playing the wrong music in the sequence.\n")
    }
    err = MusicPlay(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicPlay() called when playing already should be an no operation task with no errors.\n")
    }
    err = MusicStop(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicStop() has returned an error when it should not.\n")
    } else if !eutherpeVars.Player.Stopped {
        t.Errorf("MusicStop() seems not to be setting Player.Stopped flag accordingly.\n")
    } else if len(eutherpeVars.Player.NowPlaying.FilePath) > 0 {
        t.Errorf("MusicStop() seems not to be emptying Player.NowPlaying register accordingly.\n")
    }
    // INFO(Rafael): Simulating that all prior song selection will be replayed.
    eutherpeVars.Player.UpNextCurrentOffset = -1
    for u, currSong := range eutherpeVars.Player.UpNext {
        fmt.Printf("=== now simulating we are playing ['%s']\n", currSong.Title)
        err = MusicPlay(eutherpeVars, userData)
        if err != nil {
            t.Errorf("MusicPlay() has returned an error when it should not : '%s'.\n", err.Error())
        }
        if eutherpeVars.Player.UpNextCurrentOffset != u {
            t.Errorf("MusicPlay() seems not to be incrementing Player.UpNextCurrentOffset register accordingly.\n")
        } else if eutherpeVars.Player.NowPlaying != currSong {
            t.Errorf("MusicPlay() seems not to be following the Player.UpNext sequence.\n")
        }
        time.Sleep(30 * time.Second)
        mplayer.Stop(eutherpeVars.Player.Handle)
        time.Sleep(30 * time.Second)
    }
    if len(eutherpeVars.Player.NowPlaying.FilePath) > 0 {
        t.Errorf("MusicPlayer() after consuming all UpNext list did not clear Player.NowPlaying register.\n")
    }
    if !eutherpeVars.Player.Stopped {
        t.Errorf("MusicPlayer() after consuming all UpNext list did not changed Player.Stopped flag to true.\n")
    }
    if eutherpeVars.Player.UpNextCurrentOffset != -1 {
        t.Errorf("MusicPlayer() after consuming all UpNext list did not changed Player.UpNextCurrentOffset register to -1.\n")
    }
    if eutherpeVars.Player.Handle != nil {
        t.Errorf("MusicPlayer() after consuming all UpNext list did not set Player.Handle register to nil.\n")
    }
}
