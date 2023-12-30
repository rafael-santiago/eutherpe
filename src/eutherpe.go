package main

import (
    _ "internal/bluebraces"
    "internal/mplayer"
    "internal/dj"
    "fmt"
    _ "time"
    //"os"
)

func main() {
    /*handle, err := mplayer.Play("/mnt/vmio/06 Dharma For One.mp3")
    if err != nil {
        fmt.Println(err)
    }
    time.Sleep(10 * time.Second)
    defer mplayer.Stop(handle)*/
    /*var s mplayer.SongInfo
    s, _ = mplayer.GetSongInfo("/mnt/vmio/06 - Venus In Force.mp3")
    fmt.Println(s.Album, s.TrackNumber, s.Title, s.Artist, s.Year, s.Genre)
    s, _ = mplayer.GetSongInfo("/mnt/vmio/06 Dharma For One.mp3")
    fmt.Println(s.Album, s.TrackNumber, s.Title, s.Artist, s.Year, s.Genre)
    s, _ = mplayer.GetSongInfo("/mnt/vmio/05 - Burn the Witch.mp3")
    fmt.Println(s.Album, s.TrackNumber, s.Title, s.Artist, s.Year, s.Genre)
    s, _ = mplayer.GetSongInfo("/mnt/vmio/05 The Electric Index Eel.mp3")
    fmt.Println(s.Album, s.TrackNumber, s.Title, s.Artist, s.Year, s.Genre)*/
    collection, _ := mplayer.LoadMusicCollection("/mnt/vmio")
    for artist, albums := range collection {
        fmt.Println(artist)
        for album, songs := range albums {
            fmt.Println(" ", album)
            for _, s := range songs {
                fmt.Println("  ", s.Title, s.TrackNumber)
            }
        }
    }
    playlist := dj.Playlist{}
    playlist.Name = "PlaylistTeste.eu"
    song, _ := mplayer.GetSongInfo("/mnt/vmio/06 - Venus In Force.mp3")
    song.AlbumCover = "data"
    playlist.Add(song)
    playlist.SaveTo("blau")
    x := dj.Playlist{}
    x.LoadFrom("blau")
    fmt.Println(x.Name)
    s, _ := x.GetSongByFilePath("/mnt/vmio/06 - Venus In Force.mp3")
    fmt.Println(s)
}
