package actions

import (
    "internal/vars"
    "net/url"
    "os"
    "testing"
)

func TestTrustBluetoothDevice(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := TrustBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("TrustBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "Malformed bluetooth-trust request." {
        t.Errorf("TrustBluetoothDevice() has failed with unexpected error : '%s'.\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldBluetoothDevice, "CacildisDev")
    err = TrustBluetoothDevice(eutherpeVars, userData)
    if err != nil {
        t.Errorf("TrustBluetoothDevice() has failed when it should not.\n")
    }
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err = TrustBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("TrustBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "exit status 1" {
        t.Errorf("TrustBluetoothDevice() has failed with unexpected error : '%s'.\n", err.Error())
    }
}

func TestUntrustBluetoothDevice(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    userData := &url.Values{}
    err := UntrustBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UntrustBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "Malformed bluetooth-untrust request." {
        t.Errorf("UntrustBluetoothDevice() has failed with unexpected error : '%s'.\n", err.Error())
    }
    userData.Add(vars.EutherpePostFieldBluetoothDevice, "CacildisDev")
    err = UntrustBluetoothDevice(eutherpeVars, userData)
    if err != nil {
        t.Errorf("UntrustBluetoothDevice() has failed when it should not.\n")
    }
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err = UntrustBluetoothDevice(eutherpeVars, userData)
    if err == nil {
        t.Errorf("UntrustBluetoothDevice() has not failed when it should.\n")
    } else if err.Error() != "exit status 1" {
        t.Errorf("UntrustBluetoothDevice() has failed with unexpected error : '%s'.\n", err.Error())
    }
}



