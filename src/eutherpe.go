package main

import (
    "internal/bluebraces"
    "fmt"
    "internal/vars"
    "internal/webui"
    "os"
    "time"
    _ "internal/mplayer"
)

func main() {
    //songInfo, err1 := mplayer.GetSongInfo("soup.mp3")
    //fmt.Println(songInfo.Artist, songInfo.Album, songInfo.Title, songInfo.FilePath, len(songInfo.AlbumCover), err1)
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
    eutherpeVars.TuneUp()
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        go tryToPairWithPreviousBluetoothDevice(eutherpeVars,
                                                eutherpeVars.CachedDevices.BlueDevId)
    }
    fmt.Printf("info: Listen at %s:%s\n", eutherpeVars.HTTPd.Addr, eutherpeVars.HTTPd.Port)
    webui.RunWebUI(eutherpeVars)
    eutherpeVars.HTTPd.AuthWatchdog.Off()
    eutherpeVars.SaveSession()
    if len(eutherpeVars.HostName) > 0 {
        eutherpeVars.MDNS.GoinHome <- true
    }
}

func tryToPairWithPreviousBluetoothDevice(eutherpeVars *vars.EutherpeVars,
                                          previousDevice string) {
    // INFO(Rafael): This function allows you to pull Eutherpe's powering
    //               plug from the socket later getting back and pairing
    //               with the previous paired bluetooth output sink again.
    //               In other words, once paired and not unpaired, Eutherpe
    //               will keep on trying to pair to it. In this way you
    //               power on your Eutherpe device, power on your bluetooth
    //               and done. The (((eth)))(((e)))(((real))) becomes real ;)
    eutherpeVars.Lock()
    if eutherpeVars.CachedDevices.BlueDevId != previousDevice {
        eutherpeVars.Unlock()
        return
    }
    bluebraces.ScanDevices(3 * time.Second)
    err := bluebraces.PairDevice(previousDevice)
    if err == nil {
        err = bluebraces.ConnectDevice(previousDevice)
    }
    shouldTryAgain :=  (err != nil && eutherpeVars.CachedDevices.BlueDevId == previousDevice)
    eutherpeVars.Unlock()
    if shouldTryAgain  {
        time.Sleep(3 * time.Second)
        go tryToPairWithPreviousBluetoothDevice(eutherpeVars, previousDevice)
    }
}
