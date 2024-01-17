package actions

import (
    "internal/vars"
    "net/url"
    "testing"
)

func TestSetStorage(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.StorageDevices = append(eutherpeVars.StorageDevices, "/dev/42")
    userData := &url.Values{}
    err := SetStorage(eutherpeVars, userData)
    if err == nil {
        t.Errorf("SetStorage() did not return an error when it should.\n")
    } else if err.Error() != "Malformed storage-set request." {
        t.Errorf("SetStorage() returned an unexpected error message.\n")
    }
    userData.Add(vars.EutherpePostFieldStorageDevice, "/dev/ziriguidum!")
    err = SetStorage(eutherpeVars, userData)
    if err == nil {
        t.Errorf("SetStorage() did not return an error when it should.\n")
    } else if err.Error() != "'/dev/ziriguidum!' seems not to be a valid storage device." {
        t.Errorf("SetStorage() returned an unexpected error message.\n")
    }
    userData.Del(vars.EutherpePostFieldStorageDevice)
    userData.Add(vars.EutherpePostFieldStorageDevice, "/dev/42")
    err = SetStorage(eutherpeVars, userData)
    if err != nil {
        t.Errorf("SetStorage() returned an error when it should not.\n")
    } else if eutherpeVars.CachedDevices.MusicDevId != "/dev/42" {
        t.Errorf("SetStorage() seems not to be setting CachedDevices.MusicDevId accordingly.\n")
    }
}

