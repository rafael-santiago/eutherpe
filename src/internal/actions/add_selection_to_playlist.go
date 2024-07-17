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
    "net/url"
    "fmt"
)

func AddSelectionToPlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed addselectiontoplaylist request.")
    }
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has || len(playlist) != 1 {
        return fmt.Errorf("Malformed addselectiontoplaylist request.")
    }
    editedPlaylist := dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists)
    if editedPlaylist == nil {
        eutherpeVars.Unlock()
        err := CreatePlaylist(eutherpeVars, userData)
        if err != nil {
            eutherpeVars.Lock()
            return err
        }
        eutherpeVars.Lock()
        editedPlaylist = dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists)
        if editedPlaylist == nil {
            return fmt.Errorf("Null playlist.")
        }
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            return err
        }
        editedPlaylist.Add(song)
    }
    eutherpeVars.SavePlaylist(editedPlaylist)
    eutherpeVars.PlaylistsHTML = ""
    return nil
}
