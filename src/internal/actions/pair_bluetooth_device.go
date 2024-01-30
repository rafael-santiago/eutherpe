package actions

import (
    "internal/vars"
    "internal/bluebraces"
    "net/url"
    "flag"
    "fmt"
)

func PairBluetoothDevice(eutherpeVars *vars.EutherpeVars,
                         userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    bluetoothDevice, has := (*userData)[vars.EutherpePostFieldBluetoothDevice]
    if !has {
        return fmt.Errorf("Malformed bluetooth-pair request.")
    }
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        eutherpeVars.Unlock()
        _ = UnpairBluetoothDevice(eutherpeVars, &url.Values{})
        eutherpeVars.Lock()
    }
    err := bluebraces.PairDevice(bluetoothDevice[0], customPath)
    if err != nil {
        return err
    }
    err = bluebraces.ConnectDevice(bluetoothDevice[0], customPath)
    if err == nil {
        eutherpeVars.CachedDevices.BlueDevId = bluetoothDevice[0]
    }
    return err
}