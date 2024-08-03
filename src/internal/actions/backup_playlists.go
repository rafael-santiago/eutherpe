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
    "os"
    "path"
)

func BackupPlaylists(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if len(eutherpeVars.Playlists) == 0 {
        return fmt.Errorf("There is no playlists to backup.")
    }
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return fmt.Errorf("No storage device was set.")
    }
    playlistsDir := path.Join(eutherpeVars.CachedDevices.MusicDevId,
                              vars.EutherpeMusicDevRootDir,
                              vars.EutherpeMusicDevPlaylistsDir)
    err := os.MkdirAll(playlistsDir, 0777)
    if err != nil {
        return err
    }
    for _, playlist := range eutherpeVars.Playlists {
        err = playlist.SaveTo(path.Join(playlistsDir, playlist.Name))
        if err != nil {
            break
        }
    }
    return err
}
