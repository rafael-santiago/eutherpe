package main

import (
    _ "internal/bluebraces"
    "internal/mplayer"
    "fmt"
    "time"
    //"os"
)

func main() {
    handle, err := mplayer.Play("/mnt/vmio/06 Dharma For One.mp3")
    if err != nil {
        fmt.Println(err)
    }
    time.Sleep(10 * time.Second)
    defer mplayer.Stop(handle)
}