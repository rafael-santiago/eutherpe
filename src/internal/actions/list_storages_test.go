package actions

import (
    "internal/vars"
    "internal/storage"
    "net/url"
    "testing"
)

func TestListStorages(t *testing.T) {
    eutherpeVars := &vars.EutherpeVars{}
    eutherpeVars.CachedDevices.MusicDevId = "/dev/42"
    userData := &url.Values{}
    expectedStorages := storage.GetAllAvailableStorages()
    err := ListStorages(eutherpeVars, userData)
    if err != nil {
        t.Errorf("ListStorages() returned an error when it should not.\n")
    }
    if len(eutherpeVars.CachedDevices.MusicDevId) != 0 {
        t.Errorf("ListStorages() did not reset CachedDevices.MusicDevId field.\n")
    }
    var isEqual bool = (len(expectedStorages) == len(eutherpeVars.StorageDevices))
    if isEqual {
        var isEqual bool = true
        for x, _ := range expectedStorages {
            isEqual = (eutherpeVars.StorageDevices[x] == expectedStorages[x])
            if !isEqual {
                break
            }
        }
    }
    if !isEqual {
        t.Errorf("ListStorages() did not return the expected storages. '%v' vs. '%v'\n",
                 eutherpeVars.StorageDevices, expectedStorages)
    }
}
