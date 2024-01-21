package actions

import (
    "internal/vars"
    "internal/bluebraces"
    "net/url"
    "fmt"
    "flag"
)

func UnpairBluetoothDevice(eutherpeVars *vars.EutherpeVars,
                           _ *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    if len(eutherpeVars.CachedDevices.BlueDevId) == 0 {
        return fmt.Errorf("No device to unpair.")
    }
    _ = bluebraces.DisconnectDevice(eutherpeVars.CachedDevices.BlueDevId, customPath)
    err := bluebraces.UnpairDevice(eutherpeVars.CachedDevices.BlueDevId, customPath)
    if err == nil {
        eutherpeVars.CachedDevices.BlueDevId = ""
    }
    return err
}
