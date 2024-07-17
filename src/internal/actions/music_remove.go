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
    "net/url"
    "fmt"
)

func MusicRemove(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed music-remove request.")
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        songFilePath := GetSongFilePathFromSelectionId(selectionId)
        for n, nextSong := range eutherpeVars.Player.UpNext {
            if nextSong.FilePath == songFilePath {
                if eutherpeVars.Player.NowPlaying.FilePath == songFilePath && n == eutherpeVars.Player.UpNextCurrentOffset {
                    eutherpeVars.Unlock()
                    MusicStop(eutherpeVars, nil)
                    eutherpeVars.Lock()
                }
                eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext[:n], eutherpeVars.Player.UpNext[n+1:]...)
                break
            }
        }
        for n, nextSong := range eutherpeVars.Player.UpNextBkp {
            if nextSong.FilePath == songFilePath {
                eutherpeVars.Player.UpNextBkp = append(eutherpeVars.Player.UpNextBkp[:n], eutherpeVars.Player.UpNextBkp[n+1:]...)
                break
            }
        }
        if !eutherpeVars.Player.Stopped {
            for currOff, song := range eutherpeVars.Player.UpNext {
                if eutherpeVars.Player.NowPlaying.FilePath == song.FilePath {
                    eutherpeVars.Player.UpNextCurrentOffset = currOff
                    break
                }
            }
        }
    }
    eutherpeVars.UpNextHTML = ""
    return nil
}
