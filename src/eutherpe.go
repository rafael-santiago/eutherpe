package main

import (
    "internal/bluebraces"
    "internal/mplayer"
    _ "internal/dj"
    _ "internal/storage"
    "internal/auth"
    "fmt"
    "time"
    "internal/vars"
    "internal/webui"
    "os"
)

func main() {
    /*collection, _ := mplayer.LoadMusicCollection("/media/rs/624B-F629/")
    for artist, albums := range collection {
        fmt.Println(artist)
        for album, songs := range albums {
            fmt.Println(" ", album)
            for _, s := range songs {
                fmt.Println("  ", s.Title, s.TrackNumber)
            }
        }
    }*/
    //song, _ := mplayer.GetSongInfo("carry.mp4")
    //fmt.Println(song.Title, song.Artist, song.Album, song.TrackNumber)
    //fmt.Printf("%d %d %d %d %d %d %d %d\n", song.AlbumCover[0], song.AlbumCover[1], song.AlbumCover[2], song.AlbumCover[3], song.AlbumCover[4], song.AlbumCover[5], song.AlbumCover[6], song.AlbumCover[7])
    //os.Exit(1)
    fmt.Printf("info: Initializing bluetooth subsystem... wait...\n")
    err := bluebraces.Wear()
    if err != nil {
        fmt.Printf("panic: Unable to power on bluetooth subsystem : '%s'\n", err.Error())
        os.Exit(1)
    }
    defer bluebraces.Unwear()
    fmt.Printf("info: Bluetooth subsystem initialized!\n")
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.ConfHome = "/home/rs/src/eutherpe/src/etc/eutherpe"
    eutherpeVars.Player.RepeatAll = false
    eutherpeVars.Player.RepeatOne = false
    eutherpeVars.Player.Stopped = true
    eutherpeVars.Player.VolumeLevel = mplayer.GetVolumeLevel()
    eutherpeVars.HTTPd.URLSchema = "http"
    eutherpeVars.HTTPd.Addr = "192.168.0.133:8080"
    eutherpeVars.HTTPd.PubRoot = "/home/rs/src/eutherpe/src/web"
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/js/eutherpe.js")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/css/eutherpe.css")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Filled.otf")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Regular.otf")
    data, _ := os.ReadFile("web/html/eutherpe.html")
    eutherpeVars.HTTPd.IndexHTML = string(data)
    data, _ = os.ReadFile("web/html/error.html")
    eutherpeVars.HTTPd.ErrorHTML = string(data)
    data, _ = os.ReadFile("web/html/eutherpe-gate.html")
    eutherpeVars.HTTPd.LoginHTML = string(data)
    eutherpeVars.HTTPd.AuthWatchdog = auth.NewAuthWatchdog(time.Duration(15 * time.Minute))
    eutherpeVars.HTTPd.AuthWatchdog.On()
    eutherpeVars.RestoreSession()
    webui.RunWebUI(eutherpeVars)
    eutherpeVars.HTTPd.AuthWatchdog.Off()
    eutherpeVars.SaveSession()
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