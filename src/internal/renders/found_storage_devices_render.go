package renders

import (
    "internal/vars"
    "strings"
)

func FoundStorageDevicesRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var foundStorageDevicesHTML string = "<ul class=\"nested\">"
    for _, storageDevice := range eutherpeVars.StorageDevices {
        foundStorageDevicesHTML += "<input type=\"checkbox\" id=\"" + storageDevice + "\" class=\"StorageDevice\" onclick=\"selectSingleElement(this);\">" + storageDevice + "<br>"
    }
    foundStorageDevicesHTML += "</ul>"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleFoundStorageDevices, foundStorageDevicesHTML, -1)
}
