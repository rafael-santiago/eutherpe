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
    _ "os/exec"
    "internal/wifi"
    "net"
    "internal/mdns"
    "strings"
)

func main() {
/*
    ifaces, _ := net.Interfaces()
    for _, iface := range ifaces {
        if (iface.Flags & net.FlagLoopback) == 0 {
            addrs, _ := iface.Addrs()
            for _, addr := range addrs {
                ip, _, _ := net.ParseCIDR(addr.String())
                fmt.Println(ip)
            }
        }
    }
    os.Exit(1)
*/
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
    //cmd := exec.Command("openssl",
    //                    "req", "-new", "-newkey", "rsa:2048", "-days", "3650", "-nodes", "-x509", "-keyout", "test.priv", "-out", "test.cer", "-subj", "/CN=127.0.0.1")
    //err1 := cmd.Run()
    //fmt.Println(err1.Error())
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
    eutherpeVars.HTTPd.PubRoot = "/home/rs/src/eutherpe/src/web"
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/js/eutherpe.js")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/css/eutherpe.css")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Filled.otf")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/fonts/Sabo-Regular.otf")
    eutherpeVars.HTTPd.PubFiles = append(eutherpeVars.HTTPd.PubFiles, "/cert/eutherpe.cer")
    data, _ := os.ReadFile("web/html/eutherpe.html")
    eutherpeVars.HTTPd.IndexHTML = string(data)
    data, _ = os.ReadFile("web/html/error.html")
    eutherpeVars.HTTPd.ErrorHTML = string(data)
    data, _ = os.ReadFile("web/html/eutherpe-gate.html")
    eutherpeVars.HTTPd.LoginHTML = string(data)
    eutherpeVars.HTTPd.AuthWatchdog = auth.NewAuthWatchdog(time.Duration(15 * time.Minute))
    eutherpeVars.HTTPd.AuthWatchdog.On()
    eutherpeVars.RestoreSession()
    eutherpeVars.SetAddr()
    if eutherpeVars.WLAN.ConnSession != nil {
        defer wifi.ReleaseAddr(eutherpeVars.WLAN.Iface)
        defer wifi.Stop(eutherpeVars.WLAN.ConnSession)
        defer wifi.SetIfaceDown(eutherpeVars.WLAN.Iface)
    }
    if len(eutherpeVars.HostName) > 0 {
        eutherpeVars.MDNS.GoinHome = make(chan bool)
        eutherpeVars.MDNS.Hosts = make([]mdns.MDNSHost, 0)
        ipAddr := net.ParseIP(eutherpeVars.HTTPd.Addr)
        if strings.Index(eutherpeVars.HTTPd.Addr, ".") > - 1 {
            ipAddr = ipAddr[12:16]
        }
        eutherpeVars.MDNS.Hosts = append(eutherpeVars.MDNS.Hosts, mdns.MDNSHost { eutherpeVars.HostName, ipAddr, 300, })
        go mdns.MDNSServerStart(eutherpeVars.MDNS.Hosts, eutherpeVars.MDNS.GoinHome)
    }
    eutherpeVars.HTTPd.Port = "8080"
    fmt.Printf("info: Listen at %s:%s\n", eutherpeVars.HTTPd.Addr, eutherpeVars.HTTPd.Port)
    webui.RunWebUI(eutherpeVars)
    eutherpeVars.HTTPd.AuthWatchdog.Off()
    eutherpeVars.SaveSession()
    if len(eutherpeVars.HostName) > 0 {
        eutherpeVars.MDNS.GoinHome <- true
    }
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