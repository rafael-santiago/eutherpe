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

func SelectedStorageDeviceRender(templatedInput string, eutherpeVars *vars.EutherpeVars) string {
    var selectedStorageDeviceHTML string
    if len(eutherpeVars.CachedDevices.MusicDevId) > 0 {
        selectedStorageDeviceHTML = eutherpeVars.CachedDevices.MusicDevId
    } else {
        selectedStorageDeviceHTML = "(null)"
    }
    return strings.Replace(templatedInput, vars.EutherpeTemplateNeedleSelectedStorageDevice, selectedStorageDeviceHTML, -1)
}
