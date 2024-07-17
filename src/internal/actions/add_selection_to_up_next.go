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
    "fmt"
)

func AddSelectionToUpNext(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed addselectiontoupnext request.")
    }
    selection := ParseSelection(data[0])
    var upNextNewHead []mplayer.SongInfo
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            return err
        }
        upNextNewHead = append(upNextNewHead, song)
    }
    if eutherpeVars.Player.Stopped && eutherpeVars.Player.UpNextCurrentOffset <= 0 {
        eutherpeVars.Player.UpNext = append(upNextNewHead, eutherpeVars.Player.UpNext...)
    } else {
        if eutherpeVars.Player.UpNextCurrentOffset < 0 {
            eutherpeVars.Player.UpNextCurrentOffset = 0
        }
        upNextNewHead = append(upNextNewHead, eutherpeVars.Player.UpNext[eutherpeVars.Player.UpNextCurrentOffset+1:]...)
        eutherpeVars.Player.UpNext = append(eutherpeVars.Player.UpNext[:eutherpeVars.Player.UpNextCurrentOffset+1], upNextNewHead...)
    }
    eutherpeVars.UpNextHTML = ""
    return nil
}
