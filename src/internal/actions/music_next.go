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

func MusicNext(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    if eutherpeVars.Player.Stopped ||
       eutherpeVars.Player.Handle == nil  {
        eutherpeVars.Unlock()
        return fmt.Errorf("Not playing anything by now.")
    }
    if eutherpeVars.Player.UpNextCurrentOffset > len(eutherpeVars.Player.UpNext) - 1 {
        eutherpeVars.Player.UpNextCurrentOffset = len(eutherpeVars.Player.UpNext) - 2
    }
    eutherpeVars.Unlock()
    err := MusicStop(eutherpeVars, nil)
    if err != nil {
        return err
    }
    eutherpeVars.Lock()
    var jumpIndex int = -1
    if userData != nil {
        data, has := (*userData)[vars.EutherpePostFieldSelection]
        if has && len(data) == 1 {
            selection := ParseSelection(data[0])
            if len(selection) == 1 {
                songFilePath := GetSongFilePathFromSelectionId(selection[0])
                for u, currSong := range eutherpeVars.Player.UpNext {
                    if currSong.FilePath == songFilePath {
                        jumpIndex = u
                        break
                    }
                }
            }
        }
    }
    if jumpIndex == -1 {
        if eutherpeVars.Player.UpNextCurrentOffset < len(eutherpeVars.Player.UpNext) - 1 {
            eutherpeVars.Player.UpNextCurrentOffset++
        } else if eutherpeVars.Player.RepeatAll {
            eutherpeVars.Player.UpNextCurrentOffset = 0
        }
    } else if jumpIndex > eutherpeVars.Player.UpNextCurrentOffset {
        eutherpeVars.Player.UpNextCurrentOffset = jumpIndex
    }
    eutherpeVars.Unlock()
    return MusicPlay(eutherpeVars, nil)
}
