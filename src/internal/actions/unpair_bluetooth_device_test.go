package actions

import (
    "internal/vars"
    "net/url"
    "os"
    "testing"
)

func TestUnpairBluetoothDevice(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := UnpairBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UnpairBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "No device to unpair." {
        t.Errorf("UnpairBluetoothDevice() has failed with unexpected error : '%s'.\n", err.Error())
    }
    eutherpeVars.CachedDevices.BlueDevId = "UmXablauQualquerPoremNaoUmQualquerXablau"
    err = UnpairBluetoothDevice(eutherpeVars, userData)
    if err != nil {
        t.Errorf("UnpairBluetoothDevice() has failed when it should not.\n")
    } else if len(eutherpeVars.CachedDevices.BlueDevId) != 0 {
        t.Errorf("UnpairBluetoothDevice() seems not to be clearning cached device accordingly.\n")
    }
    eutherpeVars.CachedDevices.BlueDevId = "SeiNaoSeiNao"
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err = UnpairBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UnpairBluetoothDevice() has not failed when it should.\n")
    } else if eutherpeVars.CachedDevices.BlueDevId != "SeiNaoSeiNao" {
        t.Errorf("UnpairBluetoothDevice() seems to be clearing chached device even on error states.\n")
    }
}
