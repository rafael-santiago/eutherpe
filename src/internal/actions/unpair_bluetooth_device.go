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
        removeBluetoothDevice(&eutherpeVars.BluetoothDevices, eutherpeVars.CachedDevices.BlueDevId)
        eutherpeVars.CachedDevices.BlueDevId = ""
    }
    return err
}

func removeBluetoothDevice(blueDevs *[]bluebraces.BluetoothDevice, blueDevId string) {
    var bIndex int = -1
    for b, currBlueDev := range (*blueDevs) {
        if currBlueDev.Id == blueDevId {
            bIndex = b
            break
        }
    }
    if bIndex == -1 {
        return
    }
    (*blueDevs) = append((*blueDevs)[:bIndex], (*blueDevs)[bIndex + 1:]...)
}
