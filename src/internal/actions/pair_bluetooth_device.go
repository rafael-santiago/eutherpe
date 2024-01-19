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
    bluetoothDevice, has := (*userData)[vars.EutherpePostFieldBluetoothDevice]
    if !has {
        return fmt.Errorf("Malformed bluetooth-pair request.")
    }
    var customPath string
    if flag.Lookup("test.v") != nil {
        customPath = "../bluebraces"
    }
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        _ = UnpairBluetoothDevice(eutherpeVars, &url.Values{})
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
