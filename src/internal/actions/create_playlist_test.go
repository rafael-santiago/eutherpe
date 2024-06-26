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
    "testing"
)

func TestCreatePlaylist(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := CreatePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("CreatePlaylist() did not return an error when it should.\n")
    }
    if err.Error() != "Malformed playlist-create request." {
        t.Errorf("CreatePlaylist() returned an unexpected error.\n")
    }
    userData.Add(vars.EutherpePostFieldPlaylist, "(null)")
    err = CreatePlaylist(eutherpeVars, userData)
    if err != nil {
        t.Errorf("CreatePlaylist() has failed when it should not.\n")
    }
    err = CreatePlaylist(eutherpeVars, userData)
    if err == nil {
        t.Errorf("CreatePlaylist() has not failed when it should.\n")
    }
    if err.Error() != "Playlist '(null)' already exists." {
        t.Errorf("CreatePlaylist() has returned unexpected error.\n")
    }
}
