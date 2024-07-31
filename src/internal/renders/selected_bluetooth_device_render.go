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

func SelectedBluetoothDeviceRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var selectedBluetoothDevice string
    if len(eutherpeVars.CachedDevices.BlueDevId) > 0 {
        selectedBluetoothDevice = eutherpeVars.CachedDevices.BlueDevId
    } else {
        selectedBluetoothDevice = "(null)"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleSelectedBluetoothDevice, selectedBluetoothDevice, 1)
}
