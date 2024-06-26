//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package actions

import (
    "strings"
    "encoding/json"
)

func ParseSelection(selection string) []string {
    if len(selection) == 0 {
        return []string{}
    }
    var selectionSlice []string
    _ = json.Unmarshal([]byte(selection), &selectionSlice)
    return selectionSlice
}

func GetSongFilePathFromSelectionId(selectionId string) string {
    var s int = len(selectionId) - 1
    for ; s >= 0; s-- {
        if selectionId[s] == ':' {
            break
        }
    }
    if s == -1 {
        return ""
    }
    return selectionId[s+1:]
}

func GetArtistFromSelectionId(selectionId string) string {
    return getMetaRecordInfoFromSelectionId(selectionId, 0)
}

func GetAlbumFromSelectionId(selectionId string) string {
    return getMetaRecordInfoFromSelectionId(selectionId, 1)
}

func getMetaRecordInfoFromSelectionId(selectionId string, offset int) string {
    var startOff int
    var endOff int = len(selectionId)
    for o := 0; o <= offset; o++ {
        endOff = strings.Index(selectionId[startOff:], "/") + startOff
        if endOff == -1 {
            return ""
        }
        if o != offset {
            startOff += endOff + 1
        }
    }
    return selectionId[startOff:endOff]
}

