package actions

import (
    "internal/vars"
    "internal/mplayer"
    "net/url"
    "math/rand"
    "time"
)

func MusicShuffle(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    if eutherpeVars.Player.Shuffle {
        eutherpeVars.Player.Shuffle = false
        eutherpeVars.Player.UpNext = eutherpeVars.Player.UpNextBkp
        eutherpeVars.Player.UpNextBkp = make([]mplayer.SongInfo, 0)
    } else {
        eutherpeVars.Player.Shuffle = true
        eutherpeVars.Player.UpNextBkp = eutherpeVars.Player.UpNext
        eutherpeVars.Player.UpNext = shuffle(eutherpeVars.Player.UpNext)
    }
    return nil
}

func shuffle(playlist []mplayer.SongInfo) []mplayer.SongInfo {
    if len(playlist) <= 1 {
        return playlist
    }
    totalSongs := len(playlist)
    shufflePlaylist := make([]mplayer.SongInfo, totalSongs)
    selectedIndexes := make([]int, 0)
    t := 0
    rand.Seed(time.Now().UnixNano())
    for t < totalSongs {
        s := rand.Intn(totalSongs)
        if !has(s, selectedIndexes) {
            shufflePlaylist[t] = playlist[s]
            selectedIndexes = append(selectedIndexes, s)
            t++
        }
    }
    if isEqual(shufflePlaylist, playlist) {
        return shuffle(playlist)
    }
    return shufflePlaylist
}

func has(value int, values []int) bool {
    for _, curr_value := range values {
        if curr_value == value {
            return true
        }
    }
    return false
}

func isEqual(a, b []mplayer.SongInfo) bool {
    for i, _ := range a {
        if a[i].FilePath != b[i].FilePath {
            return false
        }
    }
    return true
}
