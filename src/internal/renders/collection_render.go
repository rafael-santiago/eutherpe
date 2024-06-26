//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package renders

import (
    "internal/vars"
    "internal/mplayer"
    "strings"
)

func CollectionRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    collectionHTML := "<ul id=\"eutherpeUL\">"
    artists := mplayer.GetArtistsFromCollection(eutherpeVars.Collection)
    for _, artist := range artists {
        collectionHTML += "<li>"
        collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "\" class=\"CollectionArtist\"><span class=\"caret\">" + artist + "</span>"
        collectionHTML += "<ul class=\"nested\">"
        albums := mplayer.GetAlbumsFromArtist(artist, eutherpeVars.Collection)
        for _, album := range albums {
            collectionHTML += "<li>"
            collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "/" + album + "\" class=\"CollectionAlbum\"><span class=\"caret\">" + album + "</span>"
            collectionHTML += "<ul class=\"nested\">"
            tracks := eutherpeVars.Collection[artist][album]
            for _, track := range tracks {
                collectionHTML += "<li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "/" + album + "/" + track.Title + ":" + track.FilePath + "\" class=\"CollectionSong\">" + track.Title + "</li>"
            }
            collectionHTML += "</ul></li>"
        }
        collectionHTML += "</ul></li>"
    }
    collectionHTML += "</ul>"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleCollection, collectionHTML, -1)
}
