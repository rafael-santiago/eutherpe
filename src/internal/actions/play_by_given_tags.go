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
    "strings"
    "strconv"
    "fmt"
)

func PlayByGivenTags(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    rawTags, has := (*userData)[vars.EutherpePostFieldTags]
    if !has {
        return fmt.Errorf("Malformed collection-playbygiventags request.")
    }
    rawAmount, has := (*userData)[vars.EutherpePostFieldAmount]
    if !has {
        return fmt.Errorf("Malformed collection-playbygiventags request.")
    }
    amount, err := strconv.Atoi(rawAmount[0])
    if err != nil {
        return err
    }
    tags := strings.Split(rawTags[0], ",")
    filePaths := make([]string, 0)
    for _, tag := range tags {
        filePathsByTag := eutherpeVars.Tags.Get(strings.Trim(tag, " "))
        for _, currFilePath := range filePathsByTag {
            if !hasFilePath(currFilePath, filePaths) {
                filePaths = append(filePaths, currFilePath)
            }
        }
    }
    songs := make([]mplayer.SongInfo, 0)
    for _, filePath := range filePaths {
        artist := GetArtistFromSelectionId(filePath)
        album := GetAlbumFromSelectionId(filePath)
        songFilePath := GetSongFilePathFromSelectionId(filePath)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, songFilePath)
        if err != nil {
            continue
        }
        songs = append(songs, song)
    }
    if len(songs) == 0 {
        return fmt.Errorf("UpNext stack underflow.")
    }
    songs = shuffle(songs)
    if !eutherpeVars.Player.Stopped {
        eutherpeVars.Unlock()
        MusicClearAll(eutherpeVars, nil)
        eutherpeVars.Lock()
    }
    if len(songs) < amount {
        amount = len(songs)
    }
    eutherpeVars.Player.UpNext = songs[:amount]
    eutherpeVars.UpNextHTML = ""
    eutherpeVars.Unlock()
    err = MusicPlay(eutherpeVars, nil)
    eutherpeVars.Lock()
    return err
}

func hasFilePath(filePath string, filePaths []string) bool {
    for _, currFilePath := range filePaths {
        if filePath == currFilePath {
            return true
        }
    }
    return false
}
