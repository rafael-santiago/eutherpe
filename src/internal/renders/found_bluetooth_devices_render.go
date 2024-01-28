package renders

import (
    "internal/vars"
    "strings"
)

func FoundBluetoothDevicesRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var foundBluetoothDevicesHTML string = "<ul class=\"nested\">"
    for _, foundDevice := range eutherpeVars.BluetoothDevices {
        foundBluetoothDevicesHTML += "<input type=\"checkbox\" id=\"" + foundDevice.Id + "\" class=\"BluetoothDevice\" onclick=\"selectSingleElement(this);\">" + foundDevice.Alias + "<br>"
    }
    foundBluetoothDevicesHTML += "</ul>"
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleFoundBluetoothDevices, foundBluetoothDevicesHTML, -1)
}
