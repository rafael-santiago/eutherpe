package main

import (
    "internal/bluebraces"
    _ "internal/mplayer"
    _ "internal/dj"
    "fmt"
    _ "time"
    "internal/vars"
    "internal/webui"
    "os"
)

func main() {
    fmt.Printf("info: Initializing bluetooth subsystem... wait...\n")
    err := bluebraces.Wear()
    if err != nil {
        fmt.Printf("panic: Unable to power on bluetooth subsystem : '%s'\n", err.Error())
        os.Exit(1)
    }
    defer bluebraces.Unwear()
    fmt.Printf("info: Bluetooth subsystem initialized!\n")
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.Player.RepeatAll = false
    eutherpeVars.Player.RepeatOne = false
    eutherpeVars.HTTPd.URLSchema = "http"
    eutherpeVars.HTTPd.Addr = "192.168.0.130:8080"
    eutherpeVars.HTTPd.PubRoot = "/root/src/eutherpe/src/web"
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/js/eutherpe.js")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/css/eutherpe.css")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Filled.otf")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Regular.otf")
    data, _ := os.ReadFile("web/html/eutherpe.html")
    eutherpeVars.HTTPd.IndexHTML = string(data)
    webui.RunWebUI(eutherpeVars)
    fmt.Println("boo!")
}

/*
func _main() {
    handle, err := mplayer.Play("/mnt/vmio/06 Dharma For One.mp3")
    if err != nil {
        fmt.Println(err)
    }
    //time.Sleep(10 * time.Second)
    handle.Wait()
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
/*    collection, _ := mplayer.LoadMusicCollection("/mnt/vmio")
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
*/