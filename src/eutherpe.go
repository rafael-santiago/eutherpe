package main

import (
    "internal/bluebraces"
    "fmt"
    "internal/vars"
    "internal/webui"
    "os"
    "time"
    "internal/mplayer"
    "strings"
)

func main() {
    //str := []rune("")
    //for r, s := range str {
    //    fmt.Println(r, s)
    //}
    //os.Exit(1)
    //songInfo, err1 := mplayer.GetSongInfo("3.m4a")
    //fmt.Println(songInfo.Artist, songInfo.Album, songInfo.TrackNumber, songInfo.Title, songInfo.FilePath, len(songInfo.AlbumCover), err1)
    //rn := []rune(songInfo.Title)
    //for r, s := range rn {
    //    fmt.Println(r, s)
    //}
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
    if len(eutherpeVars.HTTPd.Addr) == 0 ||
       strings.HasPrefix(eutherpeVars.HTTPd.Addr, "169.") {
        // TIP(Rafael): This is necessary to prevent Eutherpe listening to
        //              an invalid (or APIPA, dummy) address. It would isolate
        //              server and no one would be able to reach it.
        //              It is also useful in scenarios where you are running
        //              Eutherpe embedded and the address network is about a WLAN.
        fmt.Printf("panic: Unable to get a valid IP address.")
        os.Exit(1)
    }
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
        if err == nil {
            mplayer.SetVolume(int(eutherpeVars.Player.VolumeLevel))
        }
    }
    shouldTryAgain :=  (err != nil && eutherpeVars.CachedDevices.BlueDevId == previousDevice)
    eutherpeVars.Unlock()
    if shouldTryAgain  {
        time.Sleep(3 * time.Second)
        go tryToPairWithPreviousBluetoothDevice(eutherpeVars, previousDevice)
    }
}
