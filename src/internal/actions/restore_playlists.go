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
    "internal/storage"
    "os"
    "fmt"
    "net/url"
    "path"
)

func RestorePlaylists(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    musicDevId := eutherpeVars.CachedDevices.MusicDevId
    if len(musicDevId) == 0 {
        return fmt.Errorf("No storage device was set.")
    }
    playlistsDir := path.Join(musicDevId,
                              vars.EutherpeMusicDevRootDir,
                              vars.EutherpeMusicDevPlaylistsDir)
    _, err := os.Stat(playlistsDir)
    if err != nil {
        return fmt.Errorf("The selected storage device has no previous playlists.")
    }
    musicDevSerial := storage.GetDeviceSerialNumberByMountPoint(musicDevId)
    playlistsRootPath := path.Join(eutherpeVars.ConfHome, vars.EutherpePlaylistsHome, musicDevSerial)
    files, err := os.ReadDir(playlistsDir)
    for _, file := range files {
        fileName := file.Name()
        if file.IsDir() {
            continue
        }
        playlistBlob, err := os.ReadFile(path.Join(playlistsDir, fileName))
        if err != nil {
            return err
        }
        playlistFilePath := path.Join(playlistsRootPath, fileName)
        err = os.WriteFile(playlistFilePath, playlistBlob, 0777)
        if err != nil {
            return err
        }
    }
    err = eutherpeVars.LoadPlaylists()
    if err == nil {
        eutherpeVars.PlaylistsHTML = ""
    }
    return err
}
