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

func ShowPlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-show request.")
    }
    if dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists) == nil {
        return fmt.Errorf("Playlist '%s' not exists.", playlist[0])
    }
    eutherpeVars.RenderedPlaylist = playlist[0]
    return nil
}
