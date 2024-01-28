package renders

import (
    "internal/vars"
    "strings"
)

func SelectedStorageDeviceRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var selectedStorageDeviceHTML string
    if len(eutherpeVars.CachedDevices.MusicDevId) > 0 {
        selectedStorageDeviceHTML = eutherpeVars.CachedDevices.MusicDevId
    } else {
        selectedStorageDeviceHTML = "(null)"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleSelectedStorageDevice, selectedStorageDeviceHTML, -1)
}
