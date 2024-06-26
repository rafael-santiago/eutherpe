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
    "testing"
)

func TestRemovePlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Playlists = make([]dj.Playlist, 0)
    userData := &url.Values{}
    err := RemovePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("RemovePlaylist() has not failed when it should.\n")
    }
    if err.Error() != "Malformed playlist-remove request." {
        t.Errorf("RemovePlaylist() has failed with unexpected error: '%v'\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "(null)")
    err = RemovePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("RemovePlaylist() has not failed when it should.\n")
    }
    if err.Error() != "Playlist '(null)' not exists." {
        t.Errorf("RemovePlaylist() has returned unexpected error : '%v'\n", err.Error())
    }
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist{ Name: "(null)" })
    eutherpeVars.RenderedPlaylist = "(null)"
    err = RemovePlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("RemovePlaylist() has failed when it should not.\n")
    }
    if len(eutherpeVars.Playlists) != 0 {
        t.Errorf("RemovePlaylist() seems not being removing anything. : %v\n", eutherpeVars.Playlists)
    }
    if len(eutherpeVars.RenderedPlaylist) != 0 {
        t.Errorf("RemovePlaylist() seems not to be clearing RenderedPlaylist accordingly.\n")
    }
    eutherpeVars.Playlists = append(eutherpeVars.Playlists, dj.Playlist{ Name: "(null)" })
    eutherpeVars.RenderedPlaylist = "PraLaDeBregada"
    err = RemovePlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("RemovePlaylist() has failed when it should not.\n")
    }
    if len(eutherpeVars.Playlists) != 0 {
        t.Errorf("RemovePlaylist() seems not being removing anything. : %v\n", eutherpeVars.Playlists)
    }
    if eutherpeVars.RenderedPlaylist != "PraLaDeBregada" {
        t.Errorf("RemovePlaylist() seems to be clearing RenderedPlaylist when not needed.\n")
    }
}
