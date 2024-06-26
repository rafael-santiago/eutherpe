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
    if len(eutherpeVars.CollectionHTML) == 0 {
        eutherpeVars.CollectionHTML = renderCollectionListing(eutherpeVars.Collection)
    }
    return strings.Replace(templatedInput,
                           vars.EutherpeTemplateNeedleCollection,
                           eutherpeVars.CollectionHTML, -1)
}

func renderCollectionListing(collection mplayer.MusicCollection) string {
    collectionHTML := "<ul id=\"eutherpeUL\">"
    artists := mplayer.GetArtistsFromCollection(collection)
    for _, artist := range artists {
        collectionHTML += "<li>"
        collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "\" class=\"CollectionArtist\"><span class=\"caret\">" + artist + "</span>"
        collectionHTML += "<ul class=\"nested\">"
        albums := mplayer.GetAlbumsFromArtist(artist, collection)
        for _, album := range albums {
            collectionHTML += "<li>"
            collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "/" + album + "\" class=\"CollectionAlbum\"><span class=\"caret\">" + album + "</span>"
            collectionHTML += "<ul class=\"nested\">"
            tracks := collection[artist][album]
            for _, track := range tracks {
                collectionHTML += "<li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artist + "/" + album + "/" + track.Title + ":" + track.FilePath + "\" class=\"CollectionSong\">" + track.Title + "</li>"
            }
            collectionHTML += "</ul></li>"
        }
        collectionHTML += "</ul></li>"
    }
    collectionHTML += "</ul>"
    return collectionHTML
}
