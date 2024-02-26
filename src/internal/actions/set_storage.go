package actions

import (
    "internal/vars"
    "net/url"
    "fmt"
)

func SetStorage(eutherpeVars *vars.EutherpeVars, userData *url.Values) error {
    eutherpeVars.Lock()
    defer eutherpeVars.Unlock()
    storageDevice, has := (*userData)[vars.EutherpePostFieldStorageDevice]
    if !has {
        return fmt.Errorf("Malformed storage-set request.")
    }
    var found bool
    for _, device := range eutherpeVars.StorageDevices {
        found = (device == storageDevice[0])
    }
    if !found {
        return fmt.Errorf("'%s' seems not to be a valid storage device.", storageDevice[0])
    }
    eutherpeVars.CachedDevices.MusicDevId = storageDevice[0]
    eutherpeVars.LoadCollection()
    return nil
}
