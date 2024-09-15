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

func ReproduceSelectedOnesFromPlaylist(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-reproduceselectedones request.")
    }
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed playlist-reproduceselectedones request.")
    }
    var err error
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var playlistHasFound bool
    for _, currPlaylist := range eutherpeVars.Playlists {
        if currPlaylist.Name == playlist[0] {
            playlistHasFound = true
            break
        }
    }
    if !playlistHasFound {
        return fmt.Errorf("Playlist '%s' has not found.", playlist[0])
    }
    userData.Del(vars.EutherpePostFieldSelection)
    selection := ParseSelection(data[0])
    jsonData := "["
    for s, selectionId := range selection {
        data := split(selectionId)
        if len(data) != 3 {
            return fmt.Errorf("Malformed playlist-reproduceselectedones parameter.")
        }
        jsonData += "\"" + data[1] + ":" + data[2] + "\""
        if (s + 1) < len(selection) {
            jsonData += ","
        }
    }
    jsonData += "]"
    userData.Add(vars.EutherpePostFieldSelection, jsonData)
    if len(eutherpeVars.Player.UpNext) > 0 {
        eutherpeVars.Unlock()
        MusicClearAll(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    shouldShuffle := eutherpeVars.Player.Shuffle
    eutherpeVars.Unlock()
    err = AddSelectionToUpNext(eutherpeVars, userData)
    if err != nil {
        eutherpeVars.Lock()
        return err
    }
    if shouldShuffle {
        eutherpeVars.Lock()
        eutherpeVars.Player.UpNextBkp = eutherpeVars.Player.UpNext
        eutherpeVars.Player.UpNext = shuffle(eutherpeVars.Player.UpNextBkp)
        eutherpeVars.Unlock()
    }
    err = MusicPlay(eutherpeVars, nil)
    eutherpeVars.Lock()
    return err
}

func split(selectionId string) []string {
    items := make([]string, 0)
    startOff := 0
    for endOff, _ := range selectionId {
        if len(items) == 0 {
            if selectionId[endOff] == ':' {
                items = append(items, selectionId[startOff:endOff])
                startOff = endOff + 1
            }
        } else if selectionId[endOff] == ':' && (endOff + 1) < len(selectionId) && selectionId[endOff + 1] == '/' {
            items = append(items, selectionId[startOff:endOff])
            startOff = endOff + 1
        } else if  (endOff + 1) == len(selectionId) {
            items = append(items, selectionId[startOff:endOff+1])
        }
    }
    return items
}