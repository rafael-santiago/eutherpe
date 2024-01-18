package bluebraces

import (
    "testing"
    "os"
    "time"
)

func TestWearMustPass(t *testing.T) {
    err := Wear("./")
    if err != nil {
        t.Errorf("bluebraces.Wear() has failed : %v\n", err)
    }
}

func TestWearMustFailDueToPulseAudio(t *testing.T) {
    os.Setenv("PULSEAUDIO_MUST_FAIL", "1")
    defer os.Unsetenv("PULSEAUDIO_MUST_FAIL")
    err := Wear("./")
    if err == nil {
        t.Error("bluebraces.Wear() was expected to fail while has succeeded.")
    }
    if err.Error() != "exit status 1" {
        t.Errorf("Unexpected error: '%v'\n", err)
    }
}

func TestWearMustFailDueToBluetoohctl(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := Wear("./")
    if err == nil {
        t.Error("bluebraces.Wear() was expected to fail while has succeeded.")
    }
    if err.Error() != "exit status 1" {
        t.Errorf("Unexpected error: '%v'\n", err)
    }
}

func TestUnwearMustPass(t *testing.T) {
    err := Unwear("./")
    if err != nil {
        t.Errorf("bluebraces.Unwear() has failed : %v\n", err)
    }
}

func TestUnwearMustFailDueToPulseAudio(t *testing.T) {
    os.Setenv("PULSEAUDIO_MUST_FAIL", "1")
    defer os.Unsetenv("PULSEAUDIO_MUST_FAIL")
    err := Unwear("./")
    if err == nil {
        t.Error("bluebraces.Unwear() was expected to fail while has succeeded.")
    }
    if err.Error() != "exit status 1" {
        t.Errorf("Unexpected error: '%v'\n", err)
    }
}

func TestUnwearMustFailDueToBluetoohctl(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := Unwear("./")
    if err == nil {
        t.Error("bluebraces.Wear() was expected to fail while has succeeded.")
    }
    if err.Error() != "exit status 1" {
        t.Errorf("Unexpected error: '%v'\n", err)
    }
}

func TestScanDevicesMustPass(t *testing.T) {
    blueDevs, err := ScanDevices(time.Duration(3 * time.Second), "../bluebraces")
    if len(blueDevs) == 0 {
        t.Error("bluebraces.ScanDevices() returned no devices.")
    }
    if err != nil {
        t.Error("bluebraces.ScanDevices() has failed.\n")
    }
    expected := []BluetoothDevice {
        {"E3:91:B6:02:8C:47", "GT FUN"},
        {"B5:D0:38:C0:ED:74", "EASYWAY-BLE"},
        {"BA:BA:CA:BA:BA:CA", "PHONE-BLAU"},
        {"42:42:42:42:42:42", "zaphoid-spks"},
        {"E3:91:B6:02:8C:47", "GT FUN"},
    }
    if len(blueDevs) != len(expected) {
        t.Error("Wrong quantity of devices was returned.")
    }
    for e, exp := range expected {
        if exp != blueDevs[e] {
            t.Errorf("%v != %v\n", exp, blueDevs[e])
        }
    }
}

func TestScanDevicesMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    blueDevs, err := ScanDevices(time.Duration(3 * time.Second), "../bluebraces")
    if len(blueDevs) > 0 {
        t.Error("bluebraces.ScanDevices() returned devices.")
    }
    if err.Error() != "exit status 1" {
        t.Errorf("Unexpected error: '%v'\n", err)
    }
}

func TestPairDeviceMustPass(t *testing.T) {
    err := PairDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.PairDevice() has failed.")
    }
}

func TestPairDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := PairDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.PairDevice() has succeeded.")
    }
}

func TestUnpairDeviceMustPass(t *testing.T) {
    err := UnpairDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.UnpairDevice() has failed.")
    }
}

func TestUnpairDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := UnpairDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.UnpairDevice() has succeeded.")
    }
}

func TestConnectDeviceMustPass(t *testing.T) {
    err := ConnectDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.ConnectDevice() has failed.")
    }
}

func TestConnectDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := ConnectDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.ConnectDevice() has succeeded.")
    }
}

func TestDisconnectDeviceMustPass(t *testing.T) {
    err := DisconnectDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.DisconnectDevice() has failed.")
    }
}

func TestDisconnectDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := DisconnectDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.DisconnectDevice() has succeeded.")
    }
}

func TestTrustDeviceMustPass(t *testing.T) {
    err := TrustDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.TrustDevice() has failed.")
    }
}

func TestTrustDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := TrustDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.TrustDevice() has succeeded.")
    }
}

func TestUntrustDeviceMustPass(t *testing.T) {
    err := UntrustDevice("00:00:00:00:00:00", "../bluebraces")
    if err != nil {
        t.Error("bluebraces.UntrustDevice() has failed.")
    }
}

func TestUntrustDeviceMustFail(t *testing.T) {
    os.Setenv("BLUETOOTHCTL_MUST_FAIL", "1")
    defer os.Unsetenv("BLUETOOTHCTL_MUST_FAIL")
    err := UntrustDevice("00:00:00:00:00:00", "../bluebraces")
    if err == nil {
        t.Error("bluebrances.UntrustDevice() has succeeded.")
    }
}
