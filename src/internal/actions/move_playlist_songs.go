package actions

import (
    "internal/vars"
    "internal/dj"
    "internal/mplayer"
    "net/url"
    "fmt"
)

type moveFunc func(editedPlaylist *dj.Playlist, song mplayer.SongInfo)

func MoveUpPlaylistSongs(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    return metaMove(eutherpeVars,
                    userData,
                    func(editedPlaylist *dj.Playlist, song mplayer.SongInfo) {
                        editedPlaylist.MoveUp(song)
                    })
}

func MoveDownPlaylistSongs(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    return metaMove(eutherpeVars,
                    userData,
                    func(editedPlaylist *dj.Playlist, song mplayer.SongInfo) {
                        editedPlaylist.MoveDown(song)
                    })
}

func metaMove(eutherpeVars *vars.EutherpeVars, userData *url.Values, move moveFunc) error {
    playlist, has := (*userData)[vars.EutherpePostFieldPlaylist]
    if !has {
        return fmt.Errorf("Malformed playlist-moveup/down request.")
    }
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed playlist-moveup/down request.")
    }
    editedPlaylist := dj.GetPlaylist(playlist[0], &eutherpeVars.Playlists)
    if editedPlaylist == nil {
        return fmt.Errorf("Playlist '%s' not exists.", playlist[0])
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        artist := GetArtistFromSelectionId(selectionId)
        album := GetAlbumFromSelectionId(selectionId)
        filePath := GetSongFilePathFromSelectionId(selectionId)
        song, err := eutherpeVars.Collection.GetSongFromArtistAlbum(artist, album, filePath)
        if err != nil {
            return err
        }
        move(editedPlaylist, song)
    }
    return nil
}
