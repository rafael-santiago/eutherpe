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
    "fmt"
)

func RemoveSongsFromPlaylist(eutherpeVars *vars.EutherpeVars,
                             userData *url.Values) error {
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed playlist-removesongs request.")
    }
    selections := ParseSelection(data[0])
    for _, selection := range selections {
        data := split(selection)
        if len(data) != 3 {
            return fmt.Errorf("Malformed playlist-removesongs parameter.")
        }
        var editedPlaylist *dj.Playlist
        editedPlaylist = nil
        for p, _ := range eutherpeVars.Playlists {
            if eutherpeVars.Playlists[p].Name == data[0] {
                editedPlaylist = &eutherpeVars.Playlists[p]
                break
            }
        }
        if editedPlaylist == nil {
            return fmt.Errorf("Playlists '%s' does not exist.", data[0])
        }
        editedPlaylist.Remove(mplayer.SongInfo{ FilePath: data[2] })
    }
    eutherpeVars.PlaylistsHTML = ""
    return nil
}
