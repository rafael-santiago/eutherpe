//
// Copyright (c) 2024, Rafael Santiago
// All rights reserved.
//
// This source code is licensed under the GPLv2 license found in the
// COPYING.GPLv2 file in the root directory of Eutherpe's source tree.
//
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
