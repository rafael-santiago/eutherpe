package actions

import (
    "internal/vars"
    "internal/bluebraces"
    "flag"
    "net/url"
    "time"
    "fmt"
)

func ProbeBluetoothDevices(eutherpeVars *vars.EutherpeVars,
                           _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    blueDevs, err := bluebraces.ScanDevices(time.Duration(10 * time.Second), customPath)
    if err != nil {
        fmt.Print(err)
        return err
    }
    eutherpeVars.BluetoothDevices = blueDevs
    return nil
}
