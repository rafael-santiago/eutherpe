package actions

import (
    "internal/vars"
    "internal/storage"
    "net/url"
)

func ListStorages(eutherpeVars *vars.EutherpeVars, _ *url.Values) error {
    eutherpeVars.StorageDevices = storage.GetAllAvailableStorages()
    if len(eutherpeVars.CachedDevices.MusicDevId) == 0 {
        return nil
    }
    var found bool
    for _, device := range eutherpeVars.StorageDevices {
        found = (device == eutherpeVars.CachedDevices.MusicDevId)
        if found {
            break
        }
    }
    if !found {
        eutherpeVars.CachedDevices.MusicDevId = ""
    }
    return nil
}
