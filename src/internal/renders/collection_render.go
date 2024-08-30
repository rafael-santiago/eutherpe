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
    "fmt"
)

func CollectionRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    if len(eutherpeVars.CollectionHTML) == 0 {
        eutherpeVars.CollectionHTML = renderCollectionListing(eutherpeVars.Collection)
    }
    return strings.Replace(templatedInput,
                           vars.EutherpeTemplateNeedleCollection,
                           eutherpeVars.CollectionHTML, 1)
}

func renderCollectionListing(collection mplayer.MusicCollection) string {
    var idNr uint64
    collectionHTML := "<ul id=\"eutherpeUL\">"
    artists := mplayer.GetArtistsFromCollection(collection)
    for _, artist := range artists {
        artistId := fmt.Sprintf("%s-eutpid_%d", artist, idNr)
        idNr++
        collectionHTML += "<li>"
        collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + artistId + "\" class=\"CollectionArtist\"><span class=\"caret\">" + artist + "</span>"
        collectionHTML += "<ul class=\"nested\">"
        albums := mplayer.GetAlbumsFromArtist(artist, collection)
        for _, album := range albums {
            albumId := fmt.Sprintf("%s/%s-eutpid_%d", artistId, album, idNr)
            idNr++
            collectionHTML += "<li>"
            collectionHTML += "<input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + albumId + "\" class=\"CollectionAlbum\"><span class=\"caret\">" + album + "</span>"
            collectionHTML += "<ul class=\"nested\">"
            tracks := collection[artist][album]
            for _, track := range tracks {
                trackId := fmt.Sprintf("%s-eutpid_%d", track.Title, idNr)
                idNr++
                collectionHTML += "<li><input type=\"checkbox\" onclick=\"flush_child(this);\" id=\"" + albumId + "/" + trackId + ":" + track.FilePath + "\" class=\"CollectionSong\">" + track.Title + "</li>"
            }
            collectionHTML += "</ul></li>"
        }
        collectionHTML += "</ul></li>"
    }
    collectionHTML += "</ul>"
    return collectionHTML
}
