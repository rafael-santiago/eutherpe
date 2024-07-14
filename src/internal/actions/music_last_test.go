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
    "fmt"
    "testing"
)

func TestMusicLast(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := MusicLast(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicLast() did not return an error when it should.\n")
    } else if err.Error() != "Not playing anything by now." {

        t.Errorf("MusicLast() has returned an unexpected error.\n")
    }
    eutherpeVars.Player.Stopped = true
    err = MusicLast(eutherpeVars, userData)
    if err == nil {
        t.Errorf("MusicLast() did not return an error when it should.\n")
    } else if err.Error() != "Not playing anything by now." {
        t.Errorf("MusicNext() has returned an unexpected error.\n")
    }
    regularJohn := mplayer.SongInfo { "regular-john.mp3", "Regular John", "Queens Of The Stone Age", "Queens Of The Stone Age", "1", "1998", "", "Stoner Rock", }
    deadMen := mplayer.SongInfo { "dead_men_tell_no_tales.mp3", "Dead Men Tell No Tales", "Motorhead", "Bomber", "1", "1979", "", "Speed Metal", }
    stayClean := mplayer.SongInfo { "stay-clean.mp3", "Stay Clean", "Motorhead", "Overkill", "2", "1979", "", "Speed Metal", }
    fever := mplayer.SongInfo { "fever.mp3", "Fever", "The Cramps", "Songs The Lord Taught Us", "13", "1980", "", "Psychobilly", }
    eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext, fever, stayClean, deadMen, regularJohn)
    err = MusicPlay(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicPlay() has returned an error when it should not.\n")
    }
    if eutherpeVars.Player.NowPlaying != fever {
        t.Errorf("Player seems not to be playing the beginning of the reproduction list.\n")
    } else {
        for u := 1; u < len(eutherpeVars.Player.UpNext); u++ {
            if u < len(eutherpeVars.Player.UpNext) {
                fmt.Printf("=== now playing ['%s'] going to play ['%s']\n", eutherpeVars.Player.NowPlaying.Title,
                                                                            eutherpeVars.Player.UpNext[u].Title)
                err = MusicNext(eutherpeVars, userData)
                if err != nil {
                    t.Errorf("MusicNext() has returned an error when it should not.\n")
                }
                if eutherpeVars.Player.NowPlaying != eutherpeVars.Player.UpNext[u] {
                    t.Errorf("MusicNext() seems not to be actually playing the next song : %s != %s\n", eutherpeVars.Player.NowPlaying.Title, eutherpeVars.Player.UpNext[u].Title)
                }
            } else {
                fmt.Printf("=== now playing ['%s'] and we hit the end of the reproduction list.\n", eutherpeVars.Player.NowPlaying.Title)
                err = MusicNext(eutherpeVars, userData)
                if err != nil {
                    t.Errorf("MusicNext() did return an error when it should not.\n")
                }
            }
            time.Sleep(5 * time.Second)
        }
    }
    for u := len(eutherpeVars.Player.UpNext) - 1; u > 0; u-- {
        fmt.Printf("=== now playing ['%s'] going to play ['%s']\n", eutherpeVars.Player.NowPlaying.Title,
                                                                    eutherpeVars.Player.UpNext[u - 1].Title)
        err = MusicLast(eutherpeVars, userData)
        if err != nil {
            t.Errorf("MusicLast() has returned an error when it should not.\n")
        }
        if eutherpeVars.Player.NowPlaying != eutherpeVars.Player.UpNext[u - 1] {
            t.Errorf("MusicLast() seems not to be actually playing the last song : %s != %s\n", eutherpeVars.Player.NowPlaying.Title, eutherpeVars.Player.UpNext[u - 1].Title)
        }
        time.Sleep(5 * time.Second)
    }
    err = MusicStop(eutherpeVars, userData)
    if err != nil {
        t.Errorf("MusicStop() has returned an error when it should not : '%s'.\n", err.Error())
    }
}
