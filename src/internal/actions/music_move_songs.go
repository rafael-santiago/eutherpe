package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "fmt"
)

func MusicMoveUp(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    return metaMusicMove(eutherpeVars, userData, -1)
}

func MusicMoveDown(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    return metaMusicMove(eutherpeVars, userData, +1)
}

func metaMusicMove(eutherpeVars *vars.EutherpeVars, userData *url.Values, direction int) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    data, has := (*userData)[vars.EutherpePostFieldSelection]
    if !has {
        return fmt.Errorf("Malformed %s request.", func(d int) string {
                                                        if d < 0 {
                                                            return "music-moveup"
                                                        }
                                                        return "music-movedown"
                                                    }(direction))
    }
    selection := ParseSelection(data[0])
    for _, selectionId := range selection {
        songFilePath := GetSongFilePathFromSelectionId(selectionId)
        for _, curr_song := range eutherpeVars.Player.UpNext {
            if curr_song.FilePath == songFilePath {
                eutherpeVars.Player.UpNext = metaMoveSong(curr_song, eutherpeVars.Player.UpNext, direction)
                break
            }
        }
    }
    for off, song := range eutherpeVars.Player.UpNext {
        if song == eutherpeVars.Player.NowPlaying {
            eutherpeVars.Player.UpNextCurrentOffset = off
            break
        }
    }
    return nil
}

func metaMoveSong(song mplayer.SongInfo, songs []mplayer.SongInfo, d int) []mplayer.SongInfo {
    a := getSongIndex(song, songs)
    songsLen := len(songs)
    if a == -1 ||
       a == 0 && songsLen == 1 {
        return songs
    }
    newSongs := songs
    b := a + d
    if b == -1 || b == songsLen {
        return songs
    }
    aux := newSongs[b]
    newSongs[b] = newSongs[a]
    newSongs[a] = aux
    return newSongs
}

func getSongIndex(song mplayer.SongInfo, songs []mplayer.SongInfo) int {
    for s, curr_song := range songs {
        if curr_song == song {
            return s
        }
    }
    return -1
}
