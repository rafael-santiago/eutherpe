//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
package dj

import (
    "internal/mplayer"
    "sync"
    "fmt"
    "encoding/base64"
    "encoding/json"
    "os"
    "strings"
)

type Playlist struct {
    Name string
    songs []mplayer.SongInfo
    mtx sync.Mutex
}

func (p *Playlist) Songs() []mplayer.SongInfo {
    return p.songs
}

func (p *Playlist) Add(song mplayer.SongInfo) {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    if p.getSongIndex(song) > -1 {
        return
    }
    p.songs = append(p.songs, song)
}

func (p *Playlist) Remove(song mplayer.SongInfo) {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    s := p.getSongIndex(song)
    if s == -1 {
        return
    }
    p.songs = append(p.songs[:s], p.songs[s+1:]...)
}

func (p *Playlist) MoveUp(song mplayer.SongInfo) {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    p.metaMove(song, -1)
}

func (p *Playlist) MoveDown(song mplayer.SongInfo) {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    p.metaMove(song, +1)
}

func (p *Playlist) ClearAll() {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    p.songs = make([]mplayer.SongInfo, 0)
}

func (p *Playlist) GetSongByFilePath(filePath string) (mplayer.SongInfo, error) {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    song := mplayer.SongInfo{}
    song.FilePath = filePath
    s := p.getSongIndex(song)
    if s == -1 {
        return mplayer.SongInfo{}, fmt.Errorf("'%s' not found in playlist '%s'.", filePath, p.Name)
    }
    return p.songs[s], nil
}

func (p *Playlist) GetSongIndexByFilePath(filePath string) int {
    p.mtx.Lock()
    defer p.mtx.Unlock()
    song := mplayer.SongInfo{}
    song.FilePath = filePath
    return p.getSongIndex(song)
}

func (p *Playlist) SaveTo(filePath string) error {
    songsLen := len(p.songs)
    if songsLen == 0 {
        return fmt.Errorf("Playlist '%s' is empty.\n", p.Name)
    }
    songs := make([]mplayer.SongInfo, len(p.songs))
    copy(songs, p.songs)
    for s, _ := range songs {
        isCachedCover := strings.HasPrefix(songs[s].AlbumCover, "blob-id=")
        if !isCachedCover && len(songs[s].AlbumCover) > 0 {
            songs[s].AlbumCover = base64.StdEncoding.EncodeToString([]byte(songs[s].AlbumCover))
        }
    }
    data, err := json.Marshal(struct {
        Name string
        Songs []mplayer.SongInfo
    }{
        p.Name,
        songs,
    })
    if err != nil {
        return err
    }
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()
    _, err = file.Write(data)
    return err
}

func (p *Playlist) LoadFrom(filePath string) error {
    data, err := os.ReadFile(filePath)
    if err != nil {
        return err
    }
    var aux struct {
        Name string
        Songs []mplayer.SongInfo
    }
    json.Unmarshal(data, &aux)
    for a, _ := range aux.Songs {
        isCachedCover := strings.HasPrefix(aux.Songs[a].AlbumCover, "blob-id=")
        if !isCachedCover && len(aux.Songs[a].AlbumCover) > 0 {
            data, _ := base64.StdEncoding.DecodeString(aux.Songs[a].AlbumCover)
            aux.Songs[a].AlbumCover = string(data)
        }
    }
    p.Name = aux.Name
    p.songs = aux.Songs
    return nil
}

func GetPlaylist(playlist string, playlists *[]Playlist) *Playlist {
    for p, curr_playlist := range *playlists {
        if strings.ToLower(curr_playlist.Name) == strings.ToLower(playlist) {
            return &(*playlists)[p]
        }
    }
    return nil
}

func (p *Playlist) getSongIndex(song mplayer.SongInfo) int {
    var s int
    for s = 0; s < len(p.songs); s++ {
        if p.songs[s].FilePath == song.FilePath {
            return s
        }
    }
    return -1
}

func (p *Playlist) metaMove(song mplayer.SongInfo, d int) {
    a := p.getSongIndex(song)
    songsLen := len(p.songs)
    if a == -1 ||
       a == 0 && songsLen == 1 {
        return
    }
    b := a + d
    if b == -1 || b == songsLen {
        return
    }
    aux := p.songs[b]
    p.songs[b] = p.songs[a]
    p.songs[a] = aux
}
