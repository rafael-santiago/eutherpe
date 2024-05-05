package main

import (
    "internal/bluebraces"
    "fmt"
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
    eutherpeVars.TuneUp()
    fmt.Printf("info: Listen at %s:%s\n", eutherpeVars.HTTPd.Addr, eutherpeVars.HTTPd.Port)
    webui.RunWebUI(eutherpeVars)
    eutherpeVars.HTTPd.AuthWatchdog.Off()
    eutherpeVars.SaveSession()
    if len(eutherpeVars.HostName) > 0 {
        eutherpeVars.MDNS.GoinHome <- true
    }
}
